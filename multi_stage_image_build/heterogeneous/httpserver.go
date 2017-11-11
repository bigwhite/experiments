package main

import (
	"net/http"
	"log"
        "fmt"
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to this website!\n"))
}

func main() {
	http.HandleFunc("/", home)
        fmt.Println("Webserver start")
        fmt.Println("  -> listen on port:1111")
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}


