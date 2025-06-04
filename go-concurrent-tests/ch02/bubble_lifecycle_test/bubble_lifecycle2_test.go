package bubble_lifecycle_test

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

func TestBubbleWaitsForAllGoroutines(t *testing.T) {
	var childGoroutineFinished int32

	synctest.Test(t, func(t *testing.T) {
		t.Log("Main bubble goroutine started.")

		go func() {
			// 这个子 goroutine 会在主 bubble goroutine 之后结束
			time.Sleep(10 * time.Millisecond) // 使用合成时间，这个 sleep 会被快速跳过
			atomic.StoreInt32(&childGoroutineFinished, 1)
			t.Log("Child goroutine finished.")
		}()

		// 等待子goroutine退出
		time.Sleep(11 * time.Millisecond) // 使用合成时间，这个 sleep 可以确定在子goroutine的sleep之后返回
		t.Log("Main bubble goroutine finishing its work.")
	})

	// 当 synctest.Test 返回时，childGoroutineFinished 应该已经被设置为 1
	if atomic.LoadInt32(&childGoroutineFinished) != 1 {
		t.Fatal("synctest.Test did not wait for child goroutine to finish.")
	}
	t.Log("synctest.Test returned, child goroutine confirmed finished.")
}
