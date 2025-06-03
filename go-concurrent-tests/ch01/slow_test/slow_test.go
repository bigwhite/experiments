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

	// 为了“更可靠”，增加等待时间
	time.Sleep(200 * time.Millisecond) // 测试变慢了！

	if atomic.LoadInt32(&val) != 1 {
		t.Errorf("Async task did not complete as expected. val = %d, want 1", val)
	}
}
