package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

func main() {
	m := metrics.NewTimer()
	metrics.GetOrRegister("timer.requests", m)
	go metrics.Log(metrics.DefaultRegistry, time.Second, log.Default())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		i := rand.Intn(10)
		m.Update(time.Microsecond * time.Duration(i))
	})
	http.ListenAndServe(":8080", nil)
}
