// leak_test.go
package main_test

import (
	"context"
	"io"
	"testing"
	"testing/synctest"
)

func TestGoroutineLeakWithPipe(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		pr, pw := io.Pipe()
		defer pw.Close()

		// 这个后台goroutine在pr上阻塞读取，等待数据或EOF
		go func() {
			io.ReadAll(pr)
		}()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// 主测试goroutine错误地认为cancel()可以结束测试
		// 但实际上，后台goroutine仍在pr上阻塞
		_ = pw
		_ = ctx
	})
	// 当synctest.Test返回时，它检测到后台goroutine没有退出，
	// 于是触发panic，报告goroutine泄漏。
}
