package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/get", handleGet)
	http.HandleFunc("/api/post", handlePost)
	http.HandleFunc("/api/put", handlePut)
	http.HandleFunc("/api/delete", handleDelete)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, "index.html")
	http.ServeFile(w, r, filePath)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a GET request")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a POST request")
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a PUT request")
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received a DELETE request")
}
