package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"log"
	"runtime"
)

func main() {
	const numRecords = 1_000_000
	out := io.Discard

	for i := 0; i < 5; i++ {
		runtime.GC()
	}

	var statsBefore runtime.MemStats
	runtime.ReadMemStats(&statsBefore)

	log.Println("Starting to encode with json/v2...")

	enc := jsontext.NewEncoder(out)

	// 手动写入数组开始标记
	if err := enc.WriteToken(jsontext.BeginArray); err != nil {
		log.Fatalf("Failed to write array start: %v", err)
	}

	// 逐个编码元素
	for i := 0; i < numRecords; i++ {
		// 内存中只需要一个空结构体，几乎不占空间
		record := struct{}{}
		if err := json.MarshalEncode(enc, record); err != nil {
			log.Fatalf("v2 MarshalEncode failed for record %d: %v", i, err)
		}
	}

	// 手动写入数组结束标记
	if err := enc.WriteToken(jsontext.EndArray); err != nil {
		log.Fatalf("Failed to write array end: %v", err)
	}
	log.Println("Encode finished.")

	var statsAfter runtime.MemStats
	runtime.ReadMemStats(&statsAfter)

	allocBytes := statsAfter.TotalAlloc - statsBefore.TotalAlloc
	log.Printf("Total bytes allocated during Encode: %d bytes (%.2f KiB)", allocBytes, float64(allocBytes)/1024)
}
