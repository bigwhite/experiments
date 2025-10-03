package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// CustomTime 是一个自定义类型，用于处理非标准的 "DD/MM/YYYY HH:MM:SS" 格式
type CustomTime struct {
	time.Time
}

// 为 CustomTime 实现 UnmarshalJSON 接口
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	// 首先去除字符串的引号
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	// 定义我们期望的格式
	const layout = "02/01/2006 15:04:05"
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// UnixTime 是一个自定义类型，用于处理以秒为单位的 Unix 时间戳
type UnixTime struct {
	time.Time
}

func (ut *UnixTime) UnmarshalJSON(b []byte) error {
	// 将 JSON 数字转换为 int64
	unixSec, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	ut.Time = time.Unix(unixSec, 0)
	return nil
}

// Event 结构体包含了所有不同格式的时间字段
type Event struct {
	ISOString        time.Time  `json:"iso_string"`        // 标准库直接支持
	UnixTimestamp    UnixTime   `json:"unix_timestamp"`    // 自定义类型处理
	UnixMilliseconds int64      `json:"unix_milliseconds"` // 直接用 int64 接收
	DateOnly         string     `json:"date_only"`         // 简单情况用 string
	CustomFormat     CustomTime `json:"custom_format"`     // 自定义类型处理
}

func main() {
	jsonData := []byte(`{
		"iso_string": "2023-01-15T10:30:00.000Z",
		"unix_timestamp": 1673780200,
		"unix_milliseconds": 1673780200000,
		"date_only": "2023-01-15",
		"custom_format": "15/01/2023 10:30:00"
	}`)

	var event Event
	if err := json.Unmarshal(jsonData, &event); err != nil {
		panic(err)
	}

	fmt.Printf("ISO String:        %s\n", event.ISOString.UTC())
	fmt.Printf("Unix Timestamp:    %s\n", event.UnixTimestamp.UTC())

	// 从毫秒时间戳创建 time.Time
	msTime := time.UnixMilli(event.UnixMilliseconds)
	fmt.Printf("Unix Milliseconds: %s\n", msTime.UTC())

	fmt.Printf("Date Only:         %s\n", event.DateOnly)
	fmt.Printf("Custom Format:     %s\n", event.CustomFormat.UTC()) // 假设 custom format 也是 UTC
}
