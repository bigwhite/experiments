package writearound

import (
	"context"
	"fmt"
	"log"
	"time"

	"cachestrategysdemo/internal/cache"    // Adjust path
	"cachestrategysdemo/internal/database" // Adjust path
	"cachestrategysdemo/internal/models"   // Adjust path
)

// WriteLog writes log entry directly to DB, bypassing cache.
func WriteLog(ctx context.Context, entry *models.LogEntry, db database.Database) error {
	// 1. Write directly to DB
	log.Printf("[Write-Around Write] Writing log directly to DB (ID: %s)\n", entry.ID)
	err := db.InsertLogEntry(ctx, entry)
	if err != nil {
		return fmt.Errorf("failed to write log to DB: %w", err)
	}
	return nil
}

// GetLog retrieves log entry, using Cache-Aside for reading.
func GetLog(ctx context.Context, logID string, db database.Database, memCache cache.Cache, keyPrefix string, ttl time.Duration) (*models.LogEntry, error) {
	cacheKey := keyPrefix + logID

	// 1. Check cache (Cache-Aside read path)
	if cachedVal, found := memCache.Get(cacheKey); found {
		if entry, ok := cachedVal.(*models.LogEntry); ok {
			log.Println("[Write-Around Read] Cache Hit for log:", logID)
			return entry, nil
		}
		memCache.Delete(cacheKey) // Clean up potential bad cache entry
	}

	// 2. Cache Miss
	log.Println("[Write-Around Read] Cache Miss for log:", logID)

	// 3. Fetch from Database
	entry, err := db.GetLogByID(ctx, logID)
	if err != nil {
		return nil, fmt.Errorf("failed to get log from DB: %w", err)
	}
	if entry == nil {
		return nil, nil /* Or specific error */
	}

	// 4. Store data into cache
	memCache.Set(cacheKey, entry, ttl)
	log.Println("[Write-Around Read] Log stored in cache:", logID)

	// 5. Return data
	return entry, nil
}
