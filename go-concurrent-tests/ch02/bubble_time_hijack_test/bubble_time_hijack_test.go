package bubble_time_hijack_test

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestTimePackageBehaviorInBubble(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// 1. 验证 time.Now() 返回的是 bubble 的初始时间
		startTime := time.Now()
		expectedStartTime := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
		if !startTime.Equal(expectedStartTime) {
			t.Errorf("Initial time.Now() incorrect. Got %v, want %v", startTime, expectedStartTime)
		}
		t.Logf("Bubble start time (time.Now()): %s", startTime)

		// 2. 验证 time.Sleep() 的行为
		sleepDuration := 5 * time.Second
		beforeSleep := time.Now()

		// 启动一个 goroutine 来观察 Sleep 期间的时间
		sleepObserverChan := make(chan time.Time, 1)
		go func() {
			// 这个 goroutine 会在主 goroutine 调用 Sleep 后，但在时间实际推进前运行（如果主 goroutine Sleep 了）
			// 但由于主 goroutine 的 Sleep 几乎立即返回，这里的 time.Now() 应该还是初始时间
			sleepObserverChan <- time.Now()
		}()

		t.Logf("Before time.Sleep(%v), current bubble time: %s", sleepDuration, beforeSleep)
		time.Sleep(sleepDuration) // <<<< 这个 Sleep 会几乎立即返回，并注册一个 5s 后的定时器

		// 等待观察者 goroutine 完成
		synctest.Wait()
		timeFromObserver := <-sleepObserverChan
		if !timeFromObserver.Equal(expectedStartTime) {
			// 因为主 goroutine 的 Sleep 立刻返回，子 goroutine 执行 time.Now() 时，时间尚未推进
			t.Errorf("Observer saw time %v during main's sleep, expected bubble start time %v", timeFromObserver, expectedStartTime)
		}

		afterSleepAttempt := time.Now()

		// 只有当所有 goroutine (包括主goroutine自己)都阻塞时，时间才会推进。
		// 所以 afterSleepAttempt 应该是第一次时间推进后的时间
		if afterSleepAttempt.Equal(expectedStartTime) {
			t.Errorf("time.Now() after Sleep but before Wait incorrect. Got %v, want %v", afterSleepAttempt, expectedStartTime)
		}
		t.Logf("After time.Sleep(%v) attempt, current bubble time: %s", sleepDuration, afterSleepAttempt)

		// 3. 验证 time.NewTimer()
		timerDuration := 3 * time.Second
		timer := time.NewTimer(timerDuration)

		select {
		case <-timer.C: // 这个 case 会在合成时间推进到 3s 后被选中
			t.Logf("time.NewTimer(%v) fired as expected.", timerDuration)
			timeAfterTimer := time.Now()
			expectedTimerFireTime := afterSleepAttempt.Add(timerDuration) // 基于上一次推进后的时间
			if !timeAfterTimer.Equal(expectedTimerFireTime) {
				t.Errorf("Time after timer incorrect. Got %v, want %v", timeAfterTimer, expectedTimerFireTime)
			}
		case <-time.After(timerDuration + 1*time.Second): // 设置一个更晚的超时以防万一
			t.Fatal("time.NewTimer did not fire as expected.")
		}
	})
}
