package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3, 4, 5}

	for i, v := range m {
		i := i
		v := v
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 10)
}
