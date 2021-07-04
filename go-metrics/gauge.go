package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	g := metrics.NewGauge()
	metrics.GetOrRegister("goroutines.now", g)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				c := runtime.NumGoroutine()
				g.Update(int64(c))
				fmt.Println("goroutines now =", g.Value())
			}
		}
	}()
	http.ListenAndServe(":8080", nil)
}
