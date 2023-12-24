package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match /index.html`)
	})
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "GET /static/"`)
	})
	mux.HandleFunc("example.com/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "example.com/"`)
	})
	mux.HandleFunc("example.com/{$}", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "example.com/{$}"`)
	})
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", func(w http.ResponseWriter, req *http.Request) {
		bucket := req.PathValue("bucket")
		objectname := req.PathValue("objectname")
		fmt.Fprintln(w, `match /b/{bucket}/o/{objectname...}`+":"+"bucket="+bucket+",objectname="+objectname)
	})

	http.ListenAndServe(":8080", mux)
}
