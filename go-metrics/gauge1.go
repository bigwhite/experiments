package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	g := metrics.NewGauge()
	metrics.GetOrRegister("goroutines.now", g)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				c := runtime.NumGoroutine()
				g.Update(int64(c))
			}
		}
	}()
	http.ListenAndServe(":8080", nil)
}
