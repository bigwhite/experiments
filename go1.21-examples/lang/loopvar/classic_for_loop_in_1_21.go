package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i := 0; i < len(m); i++ {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, m[i])
		}()
	}

	time.Sleep(time.Second * 10)
}
