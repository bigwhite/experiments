package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Recv a request: %v\n", *r)
	fmt.Fprintln(w, "hello, server1.com")
}

func main() {
	// 定义命令行参数
	address := flag.String("address", "localhost:9001", "server address to listen on")

	// 解析命令行参数
	flag.Parse()

	http.HandleFunc("/", handler)
	fmt.Printf("Server is listening on %s...\n", *address)
	if err := http.ListenAndServe(*address, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
