// Package main provides ...
package main

import "fmt"

func example1() {
	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0
	fmt.Println(n0, n1)
}

func op(a, b int) int {
	return a + b
}

func example2() {
	n0, n1 := 1, 2
	n0, n1 = op(n0, n1), n0
	fmt.Println(n0, n1)
}

func main() {
	example1()
	example2()
}
