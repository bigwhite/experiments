package main

import (
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m = make(map[int]int, 100)

	for i := 0; i < 100; i++ {
		m[i] = i
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 并发写
		go func(i int) {
			for n := 0; n < 100; n++ {
				n := rand.Intn(100)
				m[n] = n
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
