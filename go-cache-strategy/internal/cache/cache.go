package cache

import (
	"sync"
	"time"
)

// Cache defines the interface for a simple cache.
type Cache interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	StopCleanup() // Method to stop background tasks if any
}

// --- In-Memory Cache Implementation ---
type cacheItem struct {
	value      interface{}
	expiration int64 // Unix Nano timestamp
}

type InMemoryCache struct {
	mu    sync.RWMutex
	items map[string]cacheItem
	stop  chan struct{} // Channel to signal cleanup goroutine to stop
	wg    sync.WaitGroup
}

func NewInMemoryCache(cleanupInterval time.Duration) *InMemoryCache {
	c := &InMemoryCache{
		items: make(map[string]cacheItem),
		stop:  make(chan struct{}),
	}
	if cleanupInterval > 0 {
		c.wg.Add(1)
		go c.runCleanup(cleanupInterval)
		// log.Println("In-Memory Cache initialized with cleanup interval:", cleanupInterval)
	} else {
		// log.Println("In-Memory Cache initialized without background cleanup.")
	}
	return c
}

func (c *InMemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	expiration := time.Now().Add(ttl).UnixNano()
	c.items[key] = cacheItem{
		value:      value,
		expiration: expiration,
	}
	// log.Printf("[Cache] Set key: %s, TTL: %v\n", key, ttl)
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[key]
	c.mu.RUnlock()

	if !found {
		return nil, false
	}

	now := time.Now().UnixNano()
	if now > item.expiration {
		// Lazy deletion on read
		c.mu.Lock()
		// Double check expiry and existence before deleting
		if currentItem, stillFound := c.items[key]; stillFound && now > currentItem.expiration {
			delete(c.items, key)
			// log.Printf("[Cache] Get key: %s - Expired (lazy delete)\n", key)
		}
		c.mu.Unlock()
		return nil, false
	}

	// log.Printf("[Cache] Get key: %s - Found\n", key)
	return item.value, true
}

func (c *InMemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
	// log.Printf("[Cache] Delete key: %s\n", key)
}

func (c *InMemoryCache) runCleanup(interval time.Duration) {
	defer c.wg.Done()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now().UnixNano()
			deletedCount := 0
			for key, item := range c.items {
				if now > item.expiration {
					delete(c.items, key)
					deletedCount++
				}
			}
			c.mu.Unlock()
			// if deletedCount > 0 {
			// 	log.Printf("Cache cleanup: Removed %d expired items.\n", deletedCount)
			// }
		case <-c.stop:
			// log.Println("Cache cleanup worker stopping.")
			return
		}
	}
}

func (c *InMemoryCache) StopCleanup() {
	// Check if the cleanup goroutine was started
	// This check prevents closing a nil channel if cleanupInterval was 0
	if c.stop != nil {
		// Prevent double close
		select {
		case <-c.stop:
			// Already closed
		default:
			close(c.stop)
		}
		c.wg.Wait() // Wait for the cleanup goroutine to finish
		// log.Println("Cache cleanup worker stopped.")
	}
}
