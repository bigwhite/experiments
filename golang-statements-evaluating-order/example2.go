// Package main provides ...
package main

import "fmt"

var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int {
	fmt.Println("calling f")
	return c
}

func g() int {
	fmt.Println("calling g")
	return a
}

func sqr(x int) int {
	fmt.Println("calling sqr")
	return x * x
}

func v() int {
	fmt.Println("calling v")
	return 1
}

func u() int {
	fmt.Println("calling u")
	return 2
}

func main() {
	fmt.Println(a, b, c)
}
