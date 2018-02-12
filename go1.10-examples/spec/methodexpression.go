package main

import "fmt"

type foo struct{}
func (foo)f() {
	fmt.Println("i am foo")
}

func main() {
	interface{f()}.f(foo{})
}
