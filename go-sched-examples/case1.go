package main

import (
	"fmt"
	"time"
)

func deadloop() {
	for {
	}
}

func main() {
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
