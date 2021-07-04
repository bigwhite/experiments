package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	m := metrics.NewMeter()
	metrics.GetOrRegister("rate.requests", m)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Mark(1)
	})
	http.ListenAndServe(":8080", nil)
}
