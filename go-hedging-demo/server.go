package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func startServer() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// 模拟 10% 的长尾延迟
		if rand.Float32() < 0.1 {
			// 长尾延迟：500ms - 1000ms
			delay := 500 + rand.Intn(500)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		} else {
			// 正常响应：10ms - 50ms
			delay := 10 + rand.Intn(40)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}

		fmt.Fprintln(w, "OK")
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(100 * time.Millisecond) // 等待服务器启动
}
