package main

import "sync"

func main() {
	const workers = 4

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
