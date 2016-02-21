package main

import "sync"

func main() {
	var ready sync.WaitGroup
	const workers = 99

	ready.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			ready.Done()
			select {}
		}()
	}

	ready.Wait()
	var i *int
	*i++
}
