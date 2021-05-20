package main

import "testing"

func foo() {
	a := 11
	p := new(int)
	*p = 12
	println("addr of a is", &a)
	println("addr that p point to is", p)
}

func bar() (*int, *int) {
	m := 21
	n := 22
	println("addr of m is", &m)
	println("addr of n is", &n)
	return &m, &n
}

func main() {
	println(int(testing.AllocsPerRun(1, foo)))
	println(int(testing.AllocsPerRun(1, func() {
		bar()
	})))
}
