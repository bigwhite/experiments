package main

import (
	"encoding/json"
	"log"
	"testing"
	"testing/slogtest"
	"time"
)

func TestChanHandlerParsing(t *testing.T) {
	var ch = make(chan []byte, 100)
	h := NewChanHandler(ch)

	results := func() []map[string]any {
		var ms []map[string]any
		ticker := time.NewTicker(time.Second)
	loop:
		for {
			select {
			case line := <-ch:
				if len(line) == 0 {
					break
				}
				var m map[string]any
				if err := json.Unmarshal(line, &m); err != nil {
					t.Fatal(err)
				}
				ms = append(ms, m)
			case <-ticker.C:
				break loop
			}
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}
}
