package main

import (
	"fmt"
	"runtime"
	"strings"
	"unique"
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

	// 使用unique包
	withUnique := make([]unique.Handle[string], numItems)
	for i := 0; i < numItems; i++ {
		withUnique[i] = unique.Make(distinctStrings[i%numDistinct])
	}

	runtime.GC() // 强制GC
	printMemUsage("With unique")

	runtime.KeepAlive(withUnique)
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
