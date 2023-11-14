package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("recv request: %#v\n", *r)
		w.Header().Set("Access-Control-Allow-Origin", "http://server1.com:8081")
		w.Write([]byte("Welcome to api/data"))
	})

	http.ListenAndServe("server2.com:8082", nil)
}
