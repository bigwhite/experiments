package main

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestCacheEntryExpires(t *testing.T) {
	synctest.Run(func() {
		count := 0
		c := NewCache(2*time.Second, func(key string) int {
			count++
			return count
		})

		// Get an entry from the cache.
		if got, want := c.Get("k"), 1; got != want {
			t.Errorf("c.Get(k) = %v, want %v", got, want)
		}

		// Verify that we get the same entry when accessing it before the expiry.
		time.Sleep(1 * time.Second)
		synctest.Wait()
		if got, want := c.Get("k"), 1; got != want {
			t.Errorf("c.Get(k) = %v, want %v", got, want)
		}

		// Wait for the entry to expire and verify that we now get a new one.
		time.Sleep(3 * time.Second)
		synctest.Wait()
		if got, want := c.Get("k"), 2; got != want {
			t.Errorf("c.Get(k) = %v, want %v", got, want)
		}
	})
}
