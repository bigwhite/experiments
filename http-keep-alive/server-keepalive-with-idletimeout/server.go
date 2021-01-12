package main

import (
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request from:", r.RemoteAddr, r.Header)
	w.Write([]byte("ok"))
}

func main() {
	var s = http.Server{
		Addr:        ":8080",
		Handler:     http.HandlerFunc(Index),
		IdleTimeout: 5 * time.Second,
	}
	s.ListenAndServe()
}
