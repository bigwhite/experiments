package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors int64 // must be accessed atomically

func handleHi(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitNum := atomic.AddInt64(&visitors, 1)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(visitNum) + "!"))
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
