package main

import (
	"cmp"
	"container/heap"
	"math/rand/v2"
	"testing"

	newheap "github.com/jba/heap" // 提案参考实现
)

// === 旧版 container/heap 所需的样板代码 ===
type OldIntHeap []int

func (h OldIntHeap) Len() int           { return len(h) }
func (h OldIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h OldIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *OldIntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *OldIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// === Benchmark 测试逻辑 ===

func BenchmarkHeapComparison(b *testing.B) {
	const size = 1000
	data := make([]int, size)
	for i := range data {
		data[i] = rand.IntN(1000000)
	}

	// 测试旧版 container/heap
	b.Run("Old_Interface_Any", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			h := &OldIntHeap{}
			for _, v := range data {
				heap.Push(h, v) // 这里会发生装箱分配
			}
			for h.Len() > 0 {
				_ = heap.Pop(h).(int) // 这里需要类型断言
			}
		}
	})

	// 测试新版 jba/heap (泛型)
	b.Run("New_Generic_V2", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			h := newheap.New(cmp.Compare[int])
			for _, v := range data {
				h.Insert(v) // 强类型插入，无装箱开销
			}
			for h.Len() > 0 {
				_ = h.TakeMin() // 直接返回 int，无需断言
			}
		}
	})
}
