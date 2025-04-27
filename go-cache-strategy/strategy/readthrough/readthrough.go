package readthrough

import (
	"context"
	"fmt"
	"log"
	"time"

	"cachestrategysdemo/internal/cache"
	"cachestrategysdemo/internal/database"
)

// LoaderFunc defines the function signature for loading data on cache miss.
type LoaderFunc func(ctx context.Context, key string) (interface{}, error)

// Cache wraps a cache instance to provide Read-Through logic.
type Cache struct {
	cache      cache.Cache // Use the cache interface
	loaderFunc LoaderFunc
	ttl        time.Duration
}

// New creates a new ReadThrough cache wrapper.
func New(cache cache.Cache, loaderFunc LoaderFunc, ttl time.Duration) *Cache {
	return &Cache{cache: cache, loaderFunc: loaderFunc, ttl: ttl}
}

// Get retrieves data, using the loader on cache miss.
func (rtc *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	// 1 & 2: Check cache
	if cachedVal, found := rtc.cache.Get(key); found {
		log.Println("[Read-Through] Cache Hit for:", key)
		return cachedVal, nil
	}

	// 4: Cache Miss - Cache calls loader
	log.Println("[Read-Through] Cache Miss for:", key)
	loadedVal, err := rtc.loaderFunc(ctx, key) // Loader fetches from DB
	if err != nil {
		return nil, fmt.Errorf("loader function failed for key %s: %w", key, err)
	}
	if loadedVal == nil {
		// Handle not found from loader (optional: cache 'not found')
		return nil, nil // Or specific error
	}

	// 5: Store loaded data into cache & return
	rtc.cache.Set(key, loadedVal, rtc.ttl)
	log.Println("[Read-Through] Loaded and stored in cache:", key)
	return loadedVal, nil
}

// Example UserLoader function (needs access to DB instance)
func NewUserLoader(db database.Database, keyPrefix string) LoaderFunc {
	return func(ctx context.Context, cacheKey string) (interface{}, error) {
		userID := cacheKey[len(keyPrefix):] // Extract ID
		// log.Println("[Read-Through Loader] Loading user from DB:", userID)
		return db.GetUser(ctx, userID)
	}
}
