package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"
)

const RequestCount = 1000

func main() {
	startServer()

	fmt.Println("开始压测普通 HTTP Client...")
	normalClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	normalLatencies := runBenchmark(normalClient)

	fmt.Println("\n开始压测 Hedged HTTP Client...")
	hedgedClient := &http.Client{
		Timeout: 2 * time.Second,
		Transport: &HedgedTransport{
			Transport:   http.DefaultTransport,
			MaxAttempts: 3,                     // 最多发送3个请求
			HedgeDelay:  80 * time.Millisecond, // P95 延迟设为触发点（我们服务器正常响应 < 50ms）
		},
	}
	hedgedLatencies := runBenchmark(hedgedClient)

	// 打印统计结果
	printStats("Normal Client", normalLatencies)
	printStats("Hedged Client", hedgedLatencies)
}

func runBenchmark(client *http.Client) []time.Duration {
	var wg sync.WaitGroup
	latencies := make([]time.Duration, RequestCount)

	for i := 0; i < RequestCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			start := time.Now()
			resp, err := client.Get("http://localhost:8080/data")
			if err != nil {
				fmt.Printf("Request failed: %v\n", err)
				return
			}
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()

			latencies[index] = time.Since(start)
		}(i)
	}

	wg.Wait()
	return latencies
}

func printStats(name string, latencies []time.Duration) {
	// 去除可能的失败请求（0值）
	valid := make([]time.Duration, 0, len(latencies))
	for _, l := range latencies {
		if l > 0 {
			valid = append(valid, l)
		}
	}

	sort.Slice(valid, func(i, j int) bool {
		return valid[i] < valid[j]
	})

	if len(valid) == 0 {
		fmt.Printf("No valid responses for %s\n", name)
		return
	}

	p50 := valid[len(valid)/2]
	p95 := valid[int(float64(len(valid))*0.95)]
	p99 := valid[int(float64(len(valid))*0.99)]

	fmt.Printf("\n=== %s 统计 ===\n", name)
	fmt.Printf("请求总数: %d\n", len(valid))
	fmt.Printf("P50 延迟: %v\n", p50)
	fmt.Printf("P95 延迟: %v\n", p95)
	fmt.Printf("P99 延迟: %v\n", p99)
}
