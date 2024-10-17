package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, server1.com")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 9001...")
	if err := http.ListenAndServe("localhost:9001", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
