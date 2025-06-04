package bubble_deadlock_test

import (
	"testing"
	"testing/synctest"
)

func TestBubbleDeadlockDetection(t *testing.T) {
	// 我们需要一个辅助函数来捕获 panic
	// 在真实的测试中，我们通常不期望 panic，但这里为了演示死锁检测
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Caught expected panic: %v", r)
			// 可以在这里进一步断言 panic 的内容是否与死锁相关
			// 例如，检查错误信息是否包含 "all goroutines are asleep - deadlock!"
			// 或 synctest 特定的死锁信息
		} else {
			t.Error("Expected a panic due to deadlock, but did not get one.")
		}
	}()

	synctest.Test(t, func(t *testing.T) {
		ch1 := make(chan int) // Bubble 内的 channel
		ch2 := make(chan int) // Bubble 内的 channel

		go func() {
			<-ch1    // 等待 ch1
			ch2 <- 1 // 然后向 ch2 发送
		}()

		// 制造一个经典的 AB-BA 死锁
		// 主 bubble goroutine 等待 ch2，而上面的 goroutine 等待 ch1
		// 且没有其他 timer 或外部事件可以打破僵局
		<-ch2    // 等待 ch2
		ch1 <- 1 // 然后向 ch1 发送
	})
}
