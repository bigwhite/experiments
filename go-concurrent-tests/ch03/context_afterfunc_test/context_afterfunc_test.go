package synctest_test

import (
	"context"
	"testing"
	"testing/synctest"
)

func TestContextAfterFunc(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// 创建一个可取消的 context
		ctx, cancel := context.WithCancel(t.Context()) // 注意这里使用了 t.Context()

		afterFuncCalled := false
		context.AfterFunc(ctx, func() {
			afterFuncCalled = true
		})

		// 验证：context 取消前，回调未执行
		synctest.Wait() // 等待所有 bubble 内 goroutine (包括可能由 AfterFunc 启动的) 阻塞
		if afterFuncCalled {
			t.Fatalf("before context is canceled: AfterFunc called prematurely")
		}

		// 取消 context
		cancel()

		// 验证：context 取消后，回调已执行
		synctest.Wait() // 再次等待，确保 AfterFunc 的 goroutine 有机会执行并完成
		if !afterFuncCalled {
			t.Fatalf("after context is canceled: AfterFunc not called")
		}
	})
}
