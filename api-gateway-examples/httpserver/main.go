package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 解析命令行参数
	port := flag.Int("p", 8080, "Port number")
	apiVersion := flag.String("v", "v1", "API version")
	apiName := flag.String("n", "example", "API name")
	flag.Parse()

	// 注册处理程序
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(*r)
		fmt.Fprintf(w, "Welcome api: localhost:%d/%s/%s\n", *port, *apiVersion, *apiName)
	})

	// 启动HTTP服务器
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Server listening on port %d\n", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
