package bubble_scope_test

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestGoroutineInBubble(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ch := make(chan time.Time, 1)

		go func() {
			// 这个 goroutine 在 bubble 内启动
			// 它将使用 bubble 的合成时间
			ch <- time.Now() // 发送的是 bubble 的初始时间
		}()

		// 等待子 goroutine 执行并发送时间
		synctest.Wait()
		// 如果不 Wait，主 goroutine 可能先退出，导致 ch 未被读取

		bubbleStartTime := <-ch
		// 验证获取到的是 bubble 的特定初始时间
		// 2000-01-01 00:00:00 +0000 UTC 的 UnixNano
		if bubbleStartTime.UnixNano() != 946684800000000000 {
			t.Errorf("Expected child goroutine to use bubble time, got %v", bubbleStartTime)
		}
		t.Logf("Time from goroutine inside bubble: %v", bubbleStartTime)
	})
}
