package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	_ "expvar"

	_ "net/http/pprof"

	"github.com/valyala/fasthttp"
)

type HelloGoHandler struct {
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "Hello, Go!")
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	go func() {
		for {
			log.Println("当前routine数量:", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()

	s := &fasthttp.Server{
		Handler: fastHTTPHandler,
	}
	s.ListenAndServe(":8081")
}
