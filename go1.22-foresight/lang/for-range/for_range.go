package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for i, v := range sl {
		wg.Add(1)
		go func() {
			fmt.Printf("%d : %d\n", i, v)
			wg.Done()
		}()
	}
	wg.Wait()
}
