package main

import (
	"encoding/json"
	"io"
	"log"
	"runtime"
)

func main() {
	const numRecords = 1_000_000
	in := make([]struct{}, numRecords)
	out := io.Discard

	// 多次 GC 以清理 sync.Pools，确保测量准确
	for i := 0; i < 5; i++ {
		runtime.GC()
	}

	var statsBefore runtime.MemStats
	runtime.ReadMemStats(&statsBefore)

	log.Println("Starting to encode with json/v1...")
	encoder := json.NewEncoder(out)
	if err := encoder.Encode(&in); err != nil {
		log.Fatalf("v1 Encode failed: %v", err)
	}
	log.Println("Encode finished.")

	var statsAfter runtime.MemStats
	runtime.ReadMemStats(&statsAfter)

	allocBytes := statsAfter.TotalAlloc - statsBefore.TotalAlloc
	log.Printf("Total bytes allocated during Encode: %d bytes (%.2f MiB)", allocBytes, float64(allocBytes)/1024/1024)
}
