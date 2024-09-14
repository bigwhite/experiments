package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	numItems    = 1000000
	stringLen   = 20
	numDistinct = 1000
)

func main() {
	// 创建一些不同的字符串
	distinctStrings := make([]string, numDistinct)
	for i := 0; i < numDistinct; i++ {
		distinctStrings[i] = strings.Repeat(string(rune('A'+i%26)), stringLen)
	}

	// 不使用unique包
	withoutUnique := make([]string, numItems)
	for i := 0; i < numItems; i++ {
		withoutUnique[i] = distinctStrings[i%numDistinct]
	}

	runtime.GC() // 强制GC以确保准确的内存使用统计
	printMemUsage("Without unique")

	runtime.KeepAlive(withoutUnique)
}

func printMemUsage(label string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s:\n", label)
	fmt.Printf("  Alloc = %v MiB\n", bToMb(m.Alloc))
	fmt.Printf("  TotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
	fmt.Printf("  Sys = %v MiB\n", bToMb(m.Sys))
	fmt.Printf("  HeapAlloc = %v MiB\n", bToMb(m.HeapAlloc))
	fmt.Printf("  HeapSys = %v MiB\n", bToMb(m.HeapSys))
	fmt.Printf("  HeapInuse = %v MiB\n", bToMb(m.HeapInuse))
	fmt.Println()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
