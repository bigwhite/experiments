package synctest_test

import (
	"testing"
	"testing/synctest"
)

func TestWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		done := false
		go func() {
			// 这个 goroutine 做的唯一事情就是设置 done 为 true 然后退出
			done = true
		}()

		// synctest.Wait() 会阻塞，直到上面那个 goroutine 执行完毕并退出
		synctest.Wait()
		if !done { // 在 Wait 返回后，done 必须为 true
			t.Fatal("goroutine did not complete setting done to true")
		}
		t.Log(done) // 根据官方示例，这里会打印 true
	})
}
