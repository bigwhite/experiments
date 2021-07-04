package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	c := metrics.NewCounter()
	metrics.GetOrRegister("total.requests", c)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.Inc(1)
	})

	http.ListenAndServe(":8080", nil)
}
