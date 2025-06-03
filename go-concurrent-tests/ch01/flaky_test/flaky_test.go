package flaky

import (
	"sync/atomic"
	"testing"
	"time"
)

// 模拟一个需要一些时间才能完成的后台任务
func performAsyncTask(val *int32) {
	go func() {
		time.Sleep(10 * time.Millisecond) // 模拟耗时操作
		atomic.StoreInt32(val, 1)
	}()
}

func TestAsyncTaskCompletion(t *testing.T) {
	var val int32
	performAsyncTask(&val)

	// 开发者“期望”后台任务在 20ms 内完成
	time.Sleep(20 * time.Millisecond)

	if atomic.LoadInt32(&val) != 1 {
		t.Errorf("Async task did not complete as expected. val = %d, want 1", val)
	}
}
