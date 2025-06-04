package synctest_test

import (
	"context"
	"testing"
	"testing/synctest"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(t.Context(), timeout)
		defer cancel() // 良好的实践，即使超时也会调用

		// 验证：超时前，context 未取消
		// 注意：这里的 Sleep 时间略小于 timeout
		time.Sleep(timeout - time.Nanosecond) // 在 bubble 内，此 Sleep 几乎立即返回
		synctest.Wait()                       // 推进合成时间到 (初始时间 + timeout - 1ns)
		if err := ctx.Err(); err != nil {
			t.Fatalf("before timeout: ctx.Err() = %v, want nil\n", err)
		}

		// 验证：超时后，context 被取消
		time.Sleep(time.Nanosecond) // 再推进 1ns，使得总时间达到 timeout
		synctest.Wait()             // 推进合成时间到 (初始时间 + timeout)
		if err := ctx.Err(); err != context.DeadlineExceeded {
			t.Fatalf("after timeout: ctx.Err() = %v, want DeadlineExceeded\n", err)
		}
	})
}
