// Package main provides ...
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		fmt.Println("Http parse form err:", err)
		return
	}
	fmt.Println("SpanId =", r.Header.Get("Span-Id"))

	time.Sleep(time.Millisecond * 201)
	w.Write([]byte("service2 ok"))
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":6602", nil)
}
