package main

import "fmt"

func Sum[T int | float64](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("%T\n", Sum(1, 2, 3.5))
	fmt.Printf("%T\n", 1+2+3.5)
}
