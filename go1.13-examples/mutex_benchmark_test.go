package mutex_test

import (
	"sync"
	"testing"
)

func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}

	return total
}

func foo() {
	var mu sync.Mutex
	mu.Lock()
	sum(10)
	mu.Unlock()
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}
}
