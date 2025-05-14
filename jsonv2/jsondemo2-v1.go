package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type EventData struct {
	EventName   string        `json:"event_name"`
	Timestamp   time.Time     `json:"timestamp,format:'2006-01-02'"`   // v2: 自定义日期格式
	PreciseTime time.Time     `json:"precise_time,format:RFC3339Nano"` // v2: RFC3339 Nano 格式
	Duration    time.Duration `json:"duration"`                        // v2 默认输出 "1h2m3s" 格式
	Timeout     time.Duration `json:"timeout,format:sec"`              // v2: 以秒为单位的数字
	OldDuration time.Duration `json:"old_duration,format:nano"`        // v2: 兼容v1的纳秒数字
}

func main() {
	fmt.Println("--- Testing Time and Duration Marshaling (v2) ---")
	event := EventData{
		EventName:   "System Update",
		Timestamp:   time.Date(2025, 5, 6, 10, 30, 0, 0, time.UTC),
		PreciseTime: time.Now(),
		Duration:    time.Hour*2 + time.Minute*15,
		Timeout:     time.Second * 90,
		OldDuration: time.Millisecond * 500,
	}

	jsonData, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println(string(jsonData))

	fmt.Println("\n--- Testing Time Unmarshaling (v2) ---")
	inputTimeJSON := `{"event_name":"Test Event", "timestamp":"2024-12-25", "precise_time":"2024-12-25T08:30:05.123456789Z", "duration":"30m", "timeout":120, "old_duration": 700000000}`
	var decodedEvent EventData
	err = json.Unmarshal([]byte(inputTimeJSON), &decodedEvent)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		fmt.Printf("Unmarshaled Event (v2 expected): %+v\n", decodedEvent)
	}
}
