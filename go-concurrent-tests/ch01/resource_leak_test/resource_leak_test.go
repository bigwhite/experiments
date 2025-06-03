package resource_leak

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func leakyGoroutineProducer(data chan int) {
	for i := 0; ; i++ { // 无限循环的生产者
		// 尝试发送数据，如果 data channel 已满且没有消费者，这里会阻塞
		// 如果消费者提前退出，这个 goroutine 将永远阻塞，造成泄露
		data <- i
		time.Sleep(1 * time.Millisecond) // 模拟生产间隔
	}
}

func TestConsumer(t *testing.T) {
	// 记录测试开始前的 goroutine 数量
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("Initial goroutines: %d\n", initialGoroutines)

	data := make(chan int, 5) // 带缓冲的 channel
	go leakyGoroutineProducer(data)

	// 消费者只消费前 10 个数据
	for i := 0; i < 10; i++ {
		val := <-data
		fmt.Printf("Consumed: %d\n", val)
	}

	// 消费者退出，但生产者 leakyGoroutineProducer 仍在运行并可能阻塞在 data <- i
	// 如果 leakyGoroutineProducer 没有合适的退出机制，它就会泄露

	// 等待一段时间，看 goroutine 数量是否恢复 (这也不是可靠的检测方法)
	time.Sleep(100 * time.Millisecond)
	finalGoroutines := runtime.NumGoroutine()
	fmt.Printf("Final goroutines: %d\n", finalGoroutines)

	if finalGoroutines > initialGoroutines {
		// 注意：这个断言本身也可能 flaky，因为 runtime 内部也可能有其他 goroutine
		// 更可靠的方式是使用专门的泄露检测工具或模式
		t.Errorf("Potential goroutine leak: initial %d, final %d", initialGoroutines, finalGoroutines)
	}
}
