package main

import (
	"runtime"
	"sync/atomic"
	"testing"
)

func BenchmarkFalseSharing(b *testing.B) {
	// 使用 GOMAXPROCS 来确定并行度，这比 NumCPU 更能反映实际调度情况
	parallelism := runtime.GOMAXPROCS(0)
	if parallelism < 2 {
		b.Skip("Skipping, need at least 2 logical CPUs to run in parallel")
	}

	b.Run("Unpadded (False Sharing)", func(b *testing.B) {
		var counters CountersUnpadded
		// 使用一个原子计数器来为每个并行goroutine分配一个唯一的、稳定的ID
		var workerIDCounter uint64
		b.RunParallel(func(pb *testing.PB) {
			// 每个goroutine在开始时获取一次ID，并在其整个生命周期中保持不变
			id := atomic.AddUint64(&workerIDCounter, 1) - 1
			counterIndex := int(id) % numCounters

			for pb.Next() {
				atomic.AddUint64(&counters.counters[counterIndex], 1)
			}
		})
	})

	b.Run("Padded (No False Sharing)", func(b *testing.B) {
		var counters CountersPadded
		var workerIDCounter uint64
		b.RunParallel(func(pb *testing.PB) {
			id := atomic.AddUint64(&workerIDCounter, 1) - 1
			counterIndex := int(id) % numCounters

			for pb.Next() {
				atomic.AddUint64(&counters.counters[counterIndex].counter, 1)
			}
		})
	})
}
