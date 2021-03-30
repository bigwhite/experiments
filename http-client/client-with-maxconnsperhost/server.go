package main

import (
	"fmt"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("receive a request from:", r.RemoteAddr, r.Header)
	time.Sleep(10 * time.Second)
	w.Write([]byte("ok"))
}

func main() {
	var s = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(Index),
	}
	s.ListenAndServe()
}
