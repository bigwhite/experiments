package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := writeToChan(ch, i)
		if ok {
			fmt.Println("write[", i, "] to channel")
			i++
			continue
		}
		fmt.Println("channel is full")
	}
}

func readFromChan(ch <-chan int) (int, bool) {
	select {
	case i := <-ch:
		return i, true

	default:
		return 0, false
	}
}

func writeToChan(ch chan<- int, i int) bool {
	select {
	case ch <- i:
		return true
	default:
		return false
	}
}

func consumer(ch <-chan int) {
	for {
		i, ok := readFromChan(ch)
		if !ok {
			fmt.Println("channel is empty")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("read[", i, "] from channel")
		if i >= 5 {
			fmt.Println("comsumer exit")
			return
		}
	}
}

func main() {
	ch := make(chan int, 5)
	go producer(ch)
	go consumer(ch)

	select {}
}
