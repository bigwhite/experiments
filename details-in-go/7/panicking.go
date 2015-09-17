package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("foo defer func invoked")
	}()
	fmt.Println("foo invoked")

	bar()
	fmt.Println("do something after bar in foo")
}

func bar() {
	defer func() {
		fmt.Println("bar defer func invoked")
	}()
	fmt.Println("bar invoked")

	zoo()
	fmt.Println("do something after zoo in bar")
}

func zoo() {
	defer func() {
		fmt.Println("zoo defer func invoked")
	}()

	fmt.Println("zoo invoked")
	panic("runtime exception")
}

func main() {
	foo()
}
