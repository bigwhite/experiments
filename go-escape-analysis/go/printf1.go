package main

import "fmt"

func foo() {
	var a int = 66666666
	var b int = 77
	fmt.Printf("a = %d\n", a)
	println("addr of a in foo =", &a)
	println("addr of b in foo =", &b)
}

func main() {
	foo()
}
