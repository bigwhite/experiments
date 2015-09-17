package main

import (
	"fmt"
	"log"
)

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
		fmt.Println("zoo defer func1 invoked")
	}()

	defer func() {
		if x := recover(); x != nil {
			log.Printf("recover panic: %v in zoo recover defer func", x)

			//log.Println("another exception occurs in zoo recover defer func")
			//panic("another exception in zoo recover defer func")
		}
	}()

	defer func() {
		fmt.Println("zoo defer func2 invoked")
	}()

	fmt.Println("zoo invoked")
	panic("zoo runtime exception")
}

func main() {
	foo()
}
