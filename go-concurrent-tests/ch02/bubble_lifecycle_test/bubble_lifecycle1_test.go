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

		// 模拟主 bubble goroutine 先做一些事情然后退出
		t.Log("Main bubble goroutine finishing its work.")
	}) // 由于子goroutine没有结束，synctest.Test 会panic

	// 当 synctest.Test 返回时，childGoroutineFinished 应该已经被设置为 1
	if atomic.LoadInt32(&childGoroutineFinished) != 1 {
		t.Fatal("synctest.Test did not wait for child goroutine to finish.")
	}
	t.Log("synctest.Test returned, child goroutine confirmed finished.")
}
