package main

import (
	"fmt"
	"sync"
	"time"

	queue "github.com/bigwhite/safe-queue/safe-queue2"
)

func main() {
	var q = queue.NewSafe()
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Second)
			q.Push(i + 1)

		}
		wg.Done()
	}()

	go func() {
	LOOP:
		for {
			select {
			case <-q.C:
				for {
					i, ok := q.Pop()
					if !ok {
						// no msg available
						continue LOOP
					}

					fmt.Printf("%d\n", i.(int))
				}
			}

		}

	}()

	wg.Wait()
}
