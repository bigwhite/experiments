package main

import "fmt"

func foo() func(int) int {
	i := []int{0: 10, 1: 11, 15: 128}
	return func(n int) int {
		n += i[0]
		return n
	}
}

func bar() {
	f := foo()
	a := f(5)
	fmt.Println(a)
}

func main() {
	bar()
	g := foo()
	b := g(6)
	fmt.Println(b)
}
