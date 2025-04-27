package writethrough

import (
	"context"
	"fmt"
	"log"
	"time"

	"cachestrategysdemo/internal/cache"    // Adjust path
	"cachestrategysdemo/internal/database" // Adjust path
	"cachestrategysdemo/internal/models"   // Adjust path
)

// UpdateUser updates user info using Write-Through strategy.
func UpdateUser(ctx context.Context, user *models.User, db database.Database, memCache cache.Cache, keyPrefix string, ttl time.Duration) error {
	cacheKey := keyPrefix + user.ID

	// Decision: Write to DB first for stronger consistency guarantee.
	log.Println("[Write-Through] Writing to database first for user:", user.ID)
	err := db.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to write to database: %w", err)
	}
	log.Println("[Write-Through] Successfully wrote to database for user:", user.ID)

	// Now write to cache (best effort after successful DB write).
	log.Println("[Write-Through] Writing to cache for user:", user.ID)
	memCache.Set(cacheKey, user, ttl)
	// Log if cache write fails, but don't return error usually
	// as the primary store succeeded. Requires monitoring.

	return nil
}
