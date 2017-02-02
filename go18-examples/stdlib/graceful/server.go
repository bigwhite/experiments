package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handle a new request:", *r)
		time.Sleep(10 * time.Second)
		log.Println("Handle the request ok!")
		io.WriteString(w, "Finished!")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-exit // wait for SIGINT
	log.Println("Shutting down server...")

	// Wait no longer than 30 seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := srv.Shutdown(ctx)

	log.Println("Server gracefully stopped:", err)
}
