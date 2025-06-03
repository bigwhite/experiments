package external_dependency

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestConcurrentAPICalls(t *testing.T) {
	// 模拟一个响应有延迟的外部 API
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond) // 模拟网络延迟和处理时间
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	var wg sync.WaitGroup
	numCalls := 5
	errors := make(chan error, numCalls)

	for i := 0; i < numCalls; i++ {
		wg.Add(1)
		go func(callNum int) {
			defer wg.Done()
			client := http.Client{Timeout: 100 * time.Millisecond} // 设置客户端超时
			resp, err := client.Get(server.URL)
			if err != nil {
				errors <- err
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				// ... 错误处理 ...
			}
		}(i)
	}

	// 如何优雅地等待所有请求完成并检查错误？
	// 如果使用 time.Sleep，需要设置多长时间？
	// 如果直接 wg.Wait()，如何处理潜在的超时导致的 goroutine 阻塞？

	// 一种可能的等待方式，但仍有缺陷
	allDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(allDone)
	}()

	select {
	case <-allDone:
		// 所有 goroutine 完成
	case <-time.After(300 * time.Millisecond): // 武断的超时时间
		t.Fatal("Test timed out waiting for API calls to complete")
	}

	close(errors)
	for err := range errors {
		if err != nil {
			t.Errorf("API call failed: %v", err)
		}
	}
}
