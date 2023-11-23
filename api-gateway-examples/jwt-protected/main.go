package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建一个基本的HTTP服务器
	mux := http.NewServeMux()

	// service handler
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("%#v\n", *req)
		w.Write([]byte("invoke protected api ok"))
	})

	// 监听28085端口
	err := http.ListenAndServe(":28085", mux)
	if err != nil {
		log.Fatal(err)
	}
}
