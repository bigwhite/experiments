package main

import (
	"container/heap"
	"fmt"
)

// 1. 必须定义一个新类型
type IntHeap []int

// 2. 必须实现标准的 5 个接口方法
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 3. Push 的参数必须是 any，内部手动断言
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// 4. Pop 的返回值必须是 any，极其容易混淆
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	// 5. 必须手动 Init
	heap.Init(h)
	// 6. 调用全局函数，而不是方法
	heap.Push(h, 3)
	// 7. Pop 出来后还得手动类型断言
	fmt.Printf("minimum: %d\n", heap.Pop(h).(int))
}
