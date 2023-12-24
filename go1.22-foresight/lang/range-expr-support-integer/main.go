package main

import "fmt"

func main() {
	n := 5
	for i := range n {
		fmt.Println(i)
	}
}
