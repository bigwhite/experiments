package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of http service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081",
		"server.crt", "server.key", nil)
}
