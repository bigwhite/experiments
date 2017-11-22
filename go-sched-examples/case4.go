package main

import (
	"fmt"
	"runtime"
	"time"
)

//go:noinline
func add(a, b int) int {
	return a + b
}

func dummy() {
	add(3, 5)
}

func deadloop() {
	for {
		dummy()
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
