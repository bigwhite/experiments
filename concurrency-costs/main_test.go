// main_test.go
package concurrency_levels

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

// Counter 是我们将要实现的各种并发计数器的通用接口
type Counter interface {
	Inc()
	Value() int64
}

// benchmark an implementation of the Counter interface
func benchmark(b *testing.B, c Counter) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc()
		}
	})
}

// --- Level 2: Mutex ---
type MutexCounter struct {
	mu      sync.Mutex
	counter int64
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *MutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func BenchmarkMutexCounter(b *testing.B) {
	benchmark(b, &MutexCounter{})
}

// --- Level 2: Atomic ---
type AtomicCounter struct {
	counter int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.counter, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.counter)
}

func BenchmarkAtomicCounter(b *testing.B) {
	benchmark(b, &AtomicCounter{})
}

// --- Level 2: CAS ---
type CasCounter struct {
	counter int64
}

func (c *CasCounter) Inc() {
	for {
		old := atomic.LoadInt64(&c.counter)
		if atomic.CompareAndSwapInt64(&c.counter, old, old+1) {
			return
		}
	}
}

func (c *CasCounter) Value() int64 {
	return atomic.LoadInt64(&c.counter)
}

func BenchmarkCasCounter(b *testing.B) {
	benchmark(b, &CasCounter{})
}

// --- Level 3: Yielding Ticket Lock ---
type YieldingTicketLockCounter struct {
	ticket  uint64
	turn    uint64
	_       [48]byte // Padding to avoid false sharing of counter
	counter int64
}

func (c *YieldingTicketLockCounter) Inc() {
	myTurn := atomic.AddUint64(&c.ticket, 1) - 1

	for atomic.LoadUint64(&c.turn) != myTurn {
		runtime.Gosched() // 主动让出，大概率触发系统调用
	}

	// 临界区
	c.counter++

	atomic.AddUint64(&c.turn, 1)
}

func (c *YieldingTicketLockCounter) Value() int64 {
	// Not thread-safe for reading while Inc is running, but sufficient for benchmark
	return c.counter
}

func BenchmarkYieldingTicketLockCounter(b *testing.B) {
	benchmark(b, &YieldingTicketLockCounter{})
}

// --- Level 4: Blocking Ticket Lock ---
type BlockingTicketLockCounter struct {
	mu      sync.Mutex
	cond    *sync.Cond
	ticket  uint64
	turn    uint64
	counter int64
}

func NewBlockingTicketLockCounter() *BlockingTicketLockCounter {
	c := &BlockingTicketLockCounter{}
	c.cond = sync.NewCond(&c.mu)
	return c
}

func (c *BlockingTicketLockCounter) Inc() {
	c.mu.Lock()
	myTurn := c.ticket
	c.ticket++

	for c.turn != myTurn {
		c.cond.Wait() // Goroutine休眠，等待被唤醒
	}
	c.mu.Unlock()

	// 在锁外执行计数器递增，以缩短临界区
	atomic.AddInt64(&c.counter, 1)

	c.mu.Lock()
	c.turn++
	c.cond.Broadcast() // 唤醒所有等待的goroutines
	c.mu.Unlock()
}

func (c *BlockingTicketLockCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func BenchmarkBlockingTicketLockCounter(b *testing.B) {
	benchmark(b, NewBlockingTicketLockCounter())
}

// --- Level 5: Catastrophic Spin Lock ---
type SpinningTicketLockCounter struct {
	ticket  uint64
	turn    uint64
	_       [48]byte
	counter int64
}

func (c *SpinningTicketLockCounter) Inc() {
	myTurn := atomic.AddUint64(&c.ticket, 1) - 1
	for atomic.LoadUint64(&c.turn) != myTurn {
	}
	c.counter++
	atomic.AddUint64(&c.turn, 1)
}

func (c *SpinningTicketLockCounter) Value() int64 {
	return c.counter
}

func BenchmarkSpinningTicketLockCounter(b *testing.B) {
	// 在测试中模拟超订场景
	// 例如，在一个8核机器上，测试时设置 b.SetParallelism(2) * runtime.NumCPU()
	// 这会让goroutine数量远超GOMAXPROCS
	//b.SetParallelism(2)
	benchmark(b, &SpinningTicketLockCounter{})
}

// --- Level 3+: Cooperative High-Cost Wait ---
type CooperativeSpinningTicketLockCounter struct {
	ticket  uint64
	turn    uint64
	_       [48]byte
	counter int64
}

func (c *CooperativeSpinningTicketLockCounter) Inc() {
	myTurn := atomic.AddUint64(&c.ticket, 1) - 1
	for atomic.LoadUint64(&c.turn) != myTurn {
		// 通过主动让出，将非协作的自旋变成了协作式的等待。
		runtime.Gosched()
	}
	c.counter++
	atomic.AddUint64(&c.turn, 1)
}

func (c *CooperativeSpinningTicketLockCounter) Value() int64 {
	return c.counter
}

func BenchmarkCooperativeSpinningTicketLockCounter(b *testing.B) {
	b.SetParallelism(2)
	benchmark(b, &CooperativeSpinningTicketLockCounter{})
}

// --- Level 1: Uncontended Atomics (Sharded) ---
const numShards = 256 // 通常是CPU核心数的几倍

type ShardedCounter struct {
	shards [numShards]struct {
		counter int64
		_       [56]byte // 填充以防止伪共享（False Sharing）
	}
}

func (c *ShardedCounter) Inc() {
	// 使用随机数模拟goroutine被分散到不同分片
	idx := rand.Intn(numShards)
	atomic.AddInt64(&c.shards[idx].counter, 1)
}

func (c *ShardedCounter) Value() int64 {
	var total int64
	for i := 0; i < numShards; i++ {
		total += atomic.LoadInt64(&c.shards[i].counter)
	}
	return total
}

func BenchmarkShardedCounter(b *testing.B) {
	benchmark(b, &ShardedCounter{})
}

// --- Level 0: Vanilla Operations (Goroutine-Local) ---
// Level 0 的 Inc() 是非同步的，所以它不直接实现Counter接口
// 我们的基准测试需要特殊设计

func BenchmarkGoroutineLocalCounter(b *testing.B) {
	var totalCounter int64

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var localCounter int64 // 每个goroutine的局部变量
		for pb.Next() {
			localCounter++ // 在局部变量上操作，无任何同步！
		}
		// 在每个goroutine结束时，将局部结果原子性地加到总数上
		atomic.AddInt64(&totalCounter, localCounter)
	})
}
