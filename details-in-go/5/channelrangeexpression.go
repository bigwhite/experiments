package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
