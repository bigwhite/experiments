package main

import "fmt"

func foo() {
	var a int = 66666666
	var b int = 77
	fmt.Printf("addr of a in bar = %p\n", &a)
	println("addr of a in bar =", &a)
	println("addr of b in bar =", &b)
}

func main() {
	foo()
}
