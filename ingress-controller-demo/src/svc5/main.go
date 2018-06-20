package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("receive a request")
	fmt.Fprintf(w, "Hello, I am svc5 for ingress-controller demo!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
