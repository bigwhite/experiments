package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/click", handleClick)
	http.HandleFunc("/api/hover", handleHover)
	http.HandleFunc("/api/submit", handleSubmit)
	http.HandleFunc("/api/search", handleSearch)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, "index.html")
	http.ServeFile(w, r, filePath)
}

func handleClick(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Button was clicked!")
}

func handleHover(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You hovered over the element!")
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	fmt.Fprintf(w, "Form submitted with message: %s", message)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	fmt.Fprintf(w, "Searching for: %s", query)
}
