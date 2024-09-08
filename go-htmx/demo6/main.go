package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/events", handleSSE)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	notificationCount := 1

	for {
		notification := fmt.Sprintf("新通知 #%d: %s", notificationCount, time.Now().Format("15:04:05"))
		fmt.Fprintf(w, "data: %s\n\n", notification)
		flusher.Flush()

		notificationCount++
		time.Sleep(3 * time.Second)

		if r.Context().Err() != nil {
			return
		}
	}
}
