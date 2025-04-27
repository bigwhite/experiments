package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// Import internal packages
	"cachestrategysdemo/internal/cache"
	"cachestrategysdemo/internal/database"
	"cachestrategysdemo/internal/models"

	// Import strategy packages
	"cachestrategysdemo/strategy/cacheaside"
	"cachestrategysdemo/strategy/readthrough"
	"cachestrategysdemo/strategy/writearound"
	"cachestrategysdemo/strategy/writebehind"
	"cachestrategysdemo/strategy/writethrough"
)

const userCacheKeyPrefix = "user:"
const userCacheTTL = 5 * time.Second // Shorter TTL for demo
const logCacheKeyPrefix = "log:"
const logCacheTTL = 1 * time.Hour

func main() {
	log.Println("Starting Cache Strategy Demo...")

	// --- Setup ---
	dbPath := "./cache_demo.db"
	_ = os.Remove(dbPath) // Clean slate

	// Initialize SQLite DB (Using interface)
	dbInstance, err := database.NewSQLiteDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer func() {
		log.Println("Closing database connection.")
		dbInstance.Close()
	}()

	// Initialize In-Memory Cache (Using interface)
	cacheInstance := cache.NewInMemoryCache(10 * time.Second) // With cleanup
	defer func() {
		log.Println("Stopping cache cleanup worker.")
		cacheInstance.StopCleanup()
	}()

	// --- Initialize Write-Behind Strategy ---
	// Needs to be initialized early so the worker starts
	wbConfig := writebehind.Config{
		Cache:     cacheInstance,
		DB:        dbInstance,
		KeyPrefix: userCacheKeyPrefix,
		TTL:       userCacheTTL,
		// Using defaults for QueueSize, BatchSize, Interval
	}
	strategyWB := writebehind.New(wbConfig)
	defer strategyWB.Stop() // Ensure worker is stopped gracefully

	ctx := context.Background()

	// --- Populate Initial Data ---
	log.Println("\n--- Populating Initial Data ---")
	initialUsers := []*models.User{
		{ID: "u1", Name: "Alice"},
		{ID: "u2", Name: "Bob"},
		{ID: "u_wt", Name: "WriteThrough Init"},
		{ID: "u_wb", Name: "WriteBehind Init"},
	}
	for _, u := range initialUsers {
		if err := dbInstance.UpdateUser(ctx, u); err != nil { // Direct DB write
			log.Printf("Failed to insert initial user %s: %v", u.ID, err)
		}
	}
	log.Println("Initial data populated.")
	time.Sleep(50 * time.Millisecond)

	// --- Demonstrate Cache-Aside ---
	fmt.Println("\n--- Testing Cache-Aside ---")
	u1_ca, err_ca1 := cacheaside.GetUser(ctx, "u1", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("1. Get u1 (Miss): %+v, Err: %v\n", u1_ca, err_ca1)
	u1_ca_hit, err_ca2 := cacheaside.GetUser(ctx, "u1", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("2. Get u1 (Hit): %+v, Err: %v\n", u1_ca_hit, err_ca2)
	fmt.Println("Waiting for Cache-Aside TTL...")
	time.Sleep(userCacheTTL + 500*time.Millisecond)
	u1_ca_exp, err_ca4 := cacheaside.GetUser(ctx, "u1", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("3. Get u1 (Expired Miss): %+v, Err: %v\n", u1_ca_exp, err_ca4)

	// --- Demonstrate Read-Through (Simulated) ---
	fmt.Println("\n--- Testing Read-Through (Simulated) ---")
	userLoader := readthrough.NewUserLoader(dbInstance, userCacheKeyPrefix)
	rtCache := readthrough.New(cacheInstance, userLoader, userCacheTTL)

	u2_rt_miss, err_rt1 := rtCache.Get(ctx, userCacheKeyPrefix+"u2")
	fmt.Printf("1. Get u2 (Miss): %+v, Err: %v\n", u2_rt_miss.(*models.User), err_rt1)
	u2_rt_hit, err_rt2 := rtCache.Get(ctx, userCacheKeyPrefix+"u2")
	fmt.Printf("2. Get u2 (Hit): %+v, Err: %v\n", u2_rt_hit.(*models.User), err_rt2)

	// --- Demonstrate Write-Through ---
	fmt.Println("\n--- Testing Write-Through ---")
	userWTUpdate := &models.User{ID: "u_wt", Name: "User WT Updated!"}
	err_wt := writethrough.UpdateUser(ctx, userWTUpdate, dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("1. Update u_wt: Err: %v\n", err_wt)
	// Verify cache immediately using Cache-Aside logic
	u_wt_read, err_wt_read := cacheaside.GetUser(ctx, "u_wt", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("2. Read u_wt via Cache-Aside (expect hit): %+v, Err: %v\n", u_wt_read, err_wt_read)

	// --- Demonstrate Write-Behind ---
	fmt.Println("\n--- Testing Write-Behind ---")
	usersWB := []*models.User{
		{ID: "u_wb", Name: "User WB Updated!"},
		{ID: "u_wb_new1", Name: "New WB 1"},
		{ID: "u_wb_new2", Name: "New WB 2"},
		{ID: "u_wb_new3", Name: "New WB 3"},
		{ID: "u_wb_new4", Name: "New WB 4"},
		{ID: "u_wb_new5", Name: "New WB 5"}, // Should trigger first batch flush
		{ID: "u_wb_new6", Name: "New WB 6"}, // Starts second batch
	}
	fmt.Println("Queueing Write-Behind updates...")
	for i, u := range usersWB {
		err := strategyWB.UpdateUser(ctx, u)
		if err != nil {
			fmt.Printf(" Error queueing WB update %d (%s): %v\n", i+1, u.ID, err)
		}
	}
	// Read immediately from cache
	u_wb_read_imm, _ := cacheaside.GetUser(ctx, "u_wb", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	u_wb_new1_read_imm, _ := cacheaside.GetUser(ctx, "u_wb_new1", dbInstance, cacheInstance, userCacheKeyPrefix, userCacheTTL)
	fmt.Printf("Read u_wb from cache immediately: %+v\n", u_wb_read_imm)
	fmt.Printf("Read u_wb_new1 from cache immediately: %+v\n", u_wb_new1_read_imm)
	fmt.Println("Waiting for Write-Behind worker to flush...")
	time.Sleep(writebehind.DefaultInterval*2 + 500*time.Millisecond) // Wait enough time

	// Verify DB after flush
	u_wb_db, _ := dbInstance.GetUser(ctx, "u_wb")
	u_wb_new6_db, _ := dbInstance.GetUser(ctx, "u_wb_new6")
	fmt.Printf("Read u_wb from DB after flush: %+v\n", u_wb_db)
	fmt.Printf("Read u_wb_new6 from DB after flush: %+v\n", u_wb_new6_db)

	// --- Demonstrate Write-Around ---
	fmt.Println("\n--- Testing Write-Around ---")
	logWA1 := &models.LogEntry{ID: "log_wa1", Message: "WA Log 1", Level: "INFO", Timestamp: time.Now()}
	err_wa_w1 := writearound.WriteLog(ctx, logWA1, dbInstance)
	fmt.Printf("1. Wrote Log1 directly to DB: Err: %v\n", err_wa_w1)
	// Read log (miss first time)
	l1_wa_miss, err_wa_r1 := writearound.GetLog(ctx, "log_wa1", dbInstance, cacheInstance, logCacheKeyPrefix, logCacheTTL)
	fmt.Printf("2. Get Log1 (Miss): %+v, Err: %v\n", l1_wa_miss, err_wa_r1)
	// Read log again (hit)
	l1_wa_hit, err_wa_r2 := writearound.GetLog(ctx, "log_wa1", dbInstance, cacheInstance, logCacheKeyPrefix, logCacheTTL)
	fmt.Printf("3. Get Log1 (Hit): %+v, Err: %v\n", l1_wa_hit, err_wa_r2)

	fmt.Println("\nDemo Finished. Cleaning up via defers...")
}
