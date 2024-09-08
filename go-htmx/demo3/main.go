package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/swap/inner", handleInner)
	http.HandleFunc("/api/swap/outer", handleOuter)
	http.HandleFunc("/api/swap/text", handleText)
	http.HandleFunc("/api/swap/before", handleBefore)
	http.HandleFunc("/api/swap/afterbegin", handleAfterBegin)
	http.HandleFunc("/api/swap/beforeend", handleBeforeEnd)
	http.HandleFunc("/api/swap/after", handleAfter)
	http.HandleFunc("/api/swap/delete", handleDelete)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	currentDir, _ := os.Getwd()
	filePath := filepath.Join(currentDir, "index.html")
	http.ServeFile(w, r, filePath)
}

func handleInner(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>This content replaced the inner HTML at %s</p>", time.Now().Format(time.RFC1123))
}

func handleOuter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div id=\"outer-content\" class=\"content-box\"><p>This div replaced the entire outer HTML at %s</p></div>", time.Now().Format(time.RFC1123))
}

func handleText(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This replaced the text content at %s. <strong>HTML tags</strong> are not parsed.", time.Now().Format(time.RFC1123))
}

func handleBefore(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p class=\"item\">This content was inserted before the target div at %s</p>", time.Now().Format(time.RFC1123))
}

func handleAfterBegin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p class=\"item\">This content was inserted at the beginning of the target div at %s</p>", time.Now().Format(time.RFC1123))
}

func handleBeforeEnd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p class=\"item\">This content was inserted at the end of the target div at %s</p>", time.Now().Format(time.RFC1123))
}

func handleAfter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p class=\"item\">This content was inserted after the target div at %s</p>", time.Now().Format(time.RFC1123))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	// For delete, we don't need to send any content back
	w.WriteHeader(http.StatusOK)
}
