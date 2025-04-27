package writebehind

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"cachestrategysdemo/internal/cache"
	"cachestrategysdemo/internal/database"
	"cachestrategysdemo/internal/models"
)

const defaultQueueSize = 100
const defaultBatchSize = 10
const DefaultInterval = 2 * time.Second

// Strategy holds the state for the Write-Behind implementation.
type Strategy struct {
	cache        cache.Cache
	db           database.Database
	updateQueue  chan *models.User
	wg           sync.WaitGroup
	stopOnce     sync.Once
	cancelCtx    context.Context
	cancelFunc   context.CancelFunc
	dbWriteMutex sync.Mutex // Simple lock for batch DB writes
	keyPrefix    string
	ttl          time.Duration
	batchSize    int
	interval     time.Duration
}

// Config holds configuration for the Write-Behind strategy.
type Config struct {
	Cache     cache.Cache
	DB        database.Database
	KeyPrefix string
	TTL       time.Duration
	QueueSize int
	BatchSize int
	Interval  time.Duration
}

// New creates and starts a new Write-Behind strategy instance.
func New(cfg Config) *Strategy {
	if cfg.QueueSize <= 0 {
		cfg.QueueSize = defaultQueueSize
	}
	if cfg.BatchSize <= 0 {
		cfg.BatchSize = defaultBatchSize
	}
	if cfg.Interval <= 0 {
		cfg.Interval = DefaultInterval
	}

	ctx, cancel := context.WithCancel(context.Background())

	s := &Strategy{
		cache:       cfg.Cache,
		db:          cfg.DB,
		updateQueue: make(chan *models.User, cfg.QueueSize),
		cancelCtx:   ctx,
		cancelFunc:  cancel,
		keyPrefix:   cfg.KeyPrefix,
		ttl:         cfg.TTL,
		batchSize:   cfg.BatchSize,
		interval:    cfg.Interval,
	}

	s.wg.Add(1)
	go s.dbWriterWorker()
	log.Println("Write-Behind worker started")
	return s
}

// UpdateUser queues a user update using Write-Behind strategy.
func (s *Strategy) UpdateUser(ctx context.Context, user *models.User) error {
	cacheKey := s.keyPrefix + user.ID
	s.cache.Set(cacheKey, user, s.ttl) // Write to cache immediately

	// Add to async queue
	select {
	case s.updateQueue <- user:
		return nil // Return success to the client immediately
	default:
		// Queue is full! Critical decision needed.
		log.Printf("[Write-Behind] Error: Update queue is full. Dropping update for user: %s\n", user.ID)
		return fmt.Errorf("update queue overflow for user %s", user.ID)
	}
}

// dbWriterWorker processes the queue and writes to DB.
func (s *Strategy) dbWriterWorker() {
	defer s.wg.Done()
	batch := make([]*models.User, 0, s.batchSize)
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	log.Println("[Write-Behind Worker] Started listening...")
	for {
		select {
		case <-s.cancelCtx.Done(): // Shutdown signal
			log.Println("[Write-Behind Worker] Shutdown signal received. Flushing final batch...")
			s.flushBatchToDB(context.Background(), batch) // Use background context
			log.Println("[Write-Behind Worker] Draining queue...")
			draining := true
			for draining {
				select {
				case user, ok := <-s.updateQueue:
					if !ok {
						draining = false
						break
					}
					batch = append(batch, user)
					if len(batch) >= s.batchSize {
						s.flushBatchToDB(context.Background(), batch)
						batch = batch[:0]
					}
				default:
					draining = false // Queue empty
				}
			}
			s.flushBatchToDB(context.Background(), batch) // Final flush
			log.Println("[Write-Behind Worker] Finished.")
			return

		case user, ok := <-s.updateQueue:
			if !ok { // Should not happen if Stop is called correctly
				log.Println("[Write-Behind Worker] Update queue closed unexpectedly.")
				s.flushBatchToDB(context.Background(), batch) // Flush remaining
				return
			}
			batch = append(batch, user)
			if len(batch) >= s.batchSize {
				// log.Printf("[Write-Behind Worker] Batch full (%d), flushing...\n", len(batch))
				s.flushBatchToDB(s.cancelCtx, batch) // Use worker context
				batch = batch[:0]
			}

		case <-ticker.C: // Write periodically
			if len(batch) > 0 {
				// log.Printf("[Write-Behind Worker] Ticker fired, flushing batch (%d)...\n", len(batch))
				s.flushBatchToDB(s.cancelCtx, batch) // Use worker context
				batch = batch[:0]
			}
		}
	}
}

// flushBatchToDB writes a batch of users to the database.
func (s *Strategy) flushBatchToDB(ctx context.Context, batch []*models.User) {
	if len(batch) == 0 {
		return
	}
	// Simple mutex to prevent concurrent batch writes (can be improved)
	s.dbWriteMutex.Lock()
	defer s.dbWriteMutex.Unlock()

	// Use the BulkUpdateUsers method from the DB interface
	err := s.db.BulkUpdateUsers(ctx, batch)
	if err != nil {
		// Error logging is handled inside BulkUpdateUsers,
		// but more sophisticated error handling (retries, dead-letter queue)
		// might be needed in production.
		log.Printf("[Write-Behind Worker] Error during bulk update flush: %v\n", err)
	}
}

// Stop gracefully shuts down the Write-Behind worker.
func (s *Strategy) Stop() {
	s.stopOnce.Do(func() {
		log.Println("[Write-Behind] Signaling worker shutdown...")
		s.cancelFunc() // Signal context cancellation
		// Closing channel could also be done, but context is cleaner for signaling
		// close(s.updateQueue)
		s.wg.Wait() // Wait for worker goroutine to finish
		log.Println("[Write-Behind] Worker stopped.")
	})
}
