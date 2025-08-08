package main

import (
	"encoding/json"
	"log"
	"runtime"
	"strings"
)

func main() {
	const numRecords = 1_000_000
	// 构造一个巨大的 JSON 数组字符串，约 2MB
	value := "[" + strings.TrimSuffix(strings.Repeat("{},", numRecords), ",") + "]"
	in := strings.NewReader(value)

	// 预分配 slice 容量，以排除 slice 自身扩容对内存测量的影响
	out := make([]struct{}, 0, numRecords)

	for i := 0; i < 5; i++ {
		runtime.GC()
	}

	var statsBefore runtime.MemStats
	runtime.ReadMemStats(&statsBefore)

	log.Println("Starting to decode with json/v1...")
	decoder := json.NewDecoder(in)
	if err := decoder.Decode(&out); err != nil {
		log.Fatalf("v1 Decode failed: %v", err)
	}
	log.Println("Decode finished.")

	var statsAfter runtime.MemStats
	runtime.ReadMemStats(&statsAfter)

	allocBytes := statsAfter.TotalAlloc - statsBefore.TotalAlloc
	log.Printf("Total bytes allocated during Decode: %d bytes (%.2f MiB)", allocBytes, float64(allocBytes)/1024/1024)
}
