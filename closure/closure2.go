package main

import (
	"fmt"
	"unsafe"
)

func foo() func(int) int {
	var a, b, c int = 11, 12, 13
	return func(n int) int {
		a += n
		b += n
		c += n
		return a + b + c
	}
}

type closure struct {
	f uintptr
	a *int
	b *int
	c *int
}

func bar() {
	f := foo()
	f(5)
	pc := *(**closure)(unsafe.Pointer(&f))
	fmt.Printf("%#v\n", *pc)
	fmt.Printf("a=%d, b=%d,c=%d\n", *pc.a, *pc.b, *pc.c)
	f(6)
	fmt.Printf("a=%d, b=%d,c=%d\n", *pc.a, *pc.b, *pc.c)
}

func main() {
	bar()
}
