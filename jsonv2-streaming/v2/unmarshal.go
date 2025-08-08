package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"errors"
	"io"
	"log"
	"runtime"
	"strings"
)

func main() {
	const numRecords = 1_000_000
	value := "[" + strings.TrimSuffix(strings.Repeat("{},", numRecords), ",") + "]"
	in := strings.NewReader(value)
	_ = make([]struct{}, 0, numRecords) // out 变量在实际应用中会用到

	for i := 0; i < 5; i++ {
		runtime.GC()
	}

	var statsBefore runtime.MemStats
	runtime.ReadMemStats(&statsBefore)

	log.Println("Starting to decode with json/v2...")

	dec := jsontext.NewDecoder(in)

	// 手动读取数组开始标记 '['
	tok, err := dec.ReadToken()
	if err != nil || tok.Kind() != '[' {
		log.Fatalf("Expected array start, got %v, err: %v", tok, err)
	}

	// 循环解码数组中的每个元素
	for dec.PeekKind() != ']' {
		var record struct{}
		if err := json.UnmarshalDecode(dec, &record); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalf("v2 UnmarshalDecode failed: %v", err)
		}
		// 在实际应用中，这里会处理 record，例如：
		// out = append(out, record)
	}
	log.Println("Decode finished.")

	var statsAfter runtime.MemStats
	runtime.ReadMemStats(&statsAfter)

	allocBytes := statsAfter.TotalAlloc - statsBefore.TotalAlloc
	log.Printf("Total bytes allocated during Decode: %d bytes (%.2f KiB)", allocBytes, float64(allocBytes)/1024)
}
