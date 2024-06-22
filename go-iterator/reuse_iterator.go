package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"hello", "world", "golang", "rust", "java"}
	itor := slices.Backward(s)
	println("first loop:\n")

	for i, x := range itor {
		fmt.Println(i, x)
		if i == 3 {
			break
		}
	}

	println("\nsecond loop:\n")

	for i, x := range itor {
		fmt.Println(i, x)
	}
}
