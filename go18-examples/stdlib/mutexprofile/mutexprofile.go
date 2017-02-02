package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
)

func main() {
	var mu sync.Mutex
	var items = make(map[int]struct{})

	runtime.SetMutexProfileFraction(5)
	for i := 0; i < 1000*1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			items[i] = struct{}{}
		}(i)
	}

	http.ListenAndServe(":8888", nil)
}
