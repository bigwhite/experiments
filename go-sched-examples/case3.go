package main

import (
	"fmt"
	"runtime"
	"time"
)

func add(a, b int) int {
	return a + b
}

func deadloop() {
	for {
		add(3, 5)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
