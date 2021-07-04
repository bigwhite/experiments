package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	s := metrics.NewExpDecaySample(1028, 0.015)
	h := metrics.NewHistogram(s)
	metrics.GetOrRegister("latency.response", h)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		i := rand.Intn(10)
		h.Update(int64(time.Microsecond * time.Duration(i)))
	})
	http.ListenAndServe(":8080", nil)
}
