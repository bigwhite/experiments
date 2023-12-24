package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{11, 12, 13, 14, 15}
	var wg sync.WaitGroup
	for i := 0; i < len(sl); i++ {
		wg.Add(1)
		go func() {
			v := sl[i]
			fmt.Printf("%d : %d\n", i, v)
			wg.Done()
		}()
	}
	wg.Wait()
}
