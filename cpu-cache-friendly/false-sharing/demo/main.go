package main

const (
	cacheLineSize = 64
	// 为了更容易观察效果，我们将计数器数量增加到与常见核心数匹配
	numCounters = 16
)

// --- 对照组 A (未填充): 计数器紧密排列，可能引发伪共享 ---
type CountersUnpadded struct {
	counters [numCounters]uint64
}

// --- 对照组 B (已填充): 通过内存填充，确保每个计数器独占一个缓存行 ---
type PaddedCounter struct {
	counter uint64
	_       [cacheLineSize - 8]byte // 填充 (64-byte cache line, 8-byte uint64)
}
type CountersPadded struct {
	counters [numCounters]PaddedCounter // 跨多个缓存行，每个缓存行一个计数器
}
