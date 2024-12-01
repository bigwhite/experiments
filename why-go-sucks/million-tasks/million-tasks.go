package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// 打印当前内存使用情况和相关信息
func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// 获取当前 goroutine 数量
	numGoroutines := runtime.NumGoroutine()

	// 获取当前线程数量
	numThreads := runtime.NumCPU() // Go runtime 不直接提供线程数量，但可以通过 NumCPU 获取逻辑处理器数量

	fmt.Printf("======>\n")
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v", m.NumGC)
	fmt.Printf("\tNumGoroutines = %v", numGoroutines)
	fmt.Printf("\tNumThreads = %v\n", numThreads)
	fmt.Printf("<======\n\n")
}

// 将字节转换为 MB
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	const signal1Goroutines = 900000
	const signal2Goroutines = 90000
	const signal3Goroutines = 10000

	// 用于接收退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 控制 goroutine 的退出
	signal1Chan := make(chan struct{})
	signal2Chan := make(chan struct{})
	signal3Chan := make(chan struct{})

	var wg sync.WaitGroup
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker.C {
			printMemoryUsage()
		}
	}()

	// 等待退出信号
	go func() {
		count := 0
		for {
			<-sigChan
			count++
			if count == 1 {
				log.Println("收到第一类goroutine退出信号")
				close(signal1Chan) // 关闭 signal1Chan，通知第一类 goroutine 退出
				continue
			}
			if count == 2 {
				log.Println("收到第二类goroutine退出信号")
				close(signal2Chan) // 关闭 signal2Chan，通知第二类 goroutine 退出
				continue
			}
			log.Println("收到第三类goroutine退出信号")
			close(signal3Chan) // 关闭 signal3Chan，通知第三类 goroutine 退出
			return
		}
	}()

	// 启动第一类 goroutine（在收到 signal1 时退出）
	log.Println("开始启动第一类goroutine...")
	for i := 0; i < signal1Goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟工作
			for {
				select {
				case <-signal1Chan:
					return
				default:
					time.Sleep(10 * time.Second) // 模拟一些工作
				}
			}
		}(i)
	}
	log.Println("启动第一类goroutine(900000) ok")

	time.Sleep(time.Second * 5)

	// 启动第二类 goroutine（在收到 signal2 时退出）
	log.Println("开始启动第二类goroutine...")
	for i := 0; i < signal2Goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟工作
			for {
				select {
				case <-signal2Chan:
					return
				default:
					time.Sleep(10 * time.Second) // 模拟一些工作
				}
			}
		}(i)
	}
	log.Println("启动第二类goroutine(90000) ok")

	time.Sleep(time.Second * 5)

	// 启动第三类goroutine（随程序退出而退出）
	log.Println("开始启动第三类goroutine...")
	for i := 0; i < signal3Goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 模拟工作
			for {
				select {
				case <-signal3Chan:
					return
				default:
					time.Sleep(10 * time.Second) // 模拟一些工作
				}
			}
		}(i)
	}
	log.Println("启动第三类goroutine(90000) ok")

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("所有 goroutine 已退出，程序结束")
}
