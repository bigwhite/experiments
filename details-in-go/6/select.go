package main

import (
	"fmt"
	"time"
)

func takeARecvChannel() chan int {
	fmt.Println("invoke takeARecvChannel")
	c := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getAStorageArr() *[5]int {
	fmt.Println("invoke getAStorageArr")
	var a [5]int
	return &a
}

func takeASendChannel() chan int {
	fmt.Println("invoke takeASendChannel")
	return make(chan int)
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func main() {
	select {
	//recv channels
	case (getAStorageArr())[0] = <-takeARecvChannel():
		fmt.Println("recv something from a recv channel")

		//send channels
	case takeASendChannel() <- getANumToChannel():
		fmt.Println("send something to a send channel")
	}
}
