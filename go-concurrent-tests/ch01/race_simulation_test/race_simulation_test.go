package race_simulation

import (
	"sync"
	"testing"
	"time"
)

// 目标：测试在特定交错下，共享资源是否会被破坏
var sharedResource int
var mu sync.Mutex

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	// 模拟一些操作
	current := sharedResource
	time.Sleep(1 * time.Millisecond) // 期望在这里发生上下文切换
	sharedResource = current + 1
	mu.Unlock()
}

func TestSpecificInterleaving(t *testing.T) {
	var wg sync.WaitGroup
	sharedResource = 0

	// 尝试通过启动顺序和微小的 sleep 来“诱导”特定的交错
	// 但这极不可靠
	wg.Add(1)
	go worker(1, &wg)

	time.Sleep(1 * time.Microsecond) // 试图让 worker1 先获得锁并执行一部分

	wg.Add(1)
	go worker(2, &wg)

	// 理想情况下，我们希望测试能稳定地复现因交错导致的 sharedResource != 2 的情况
	// 但实际上很难通过 time.Sleep 精确控制
	mu.Lock()
	if sharedResource != 2 {
		t.Errorf("Race condition possibly occurred or logic error, sharedResource = %d, want 2", sharedResource)
	}
	mu.Unlock()

	wg.Wait()
}
