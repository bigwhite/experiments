package main

import (
	"fmt"
	"plugin"
)

func init() {
	fmt.Println("init in main program")
}

func main() {
	fmt.Println("before opening the foo.so")
	p, _ := plugin.Open("foo.so")
	fmt.Println("after opening the foo.so")
	f, _ := p.Lookup("Foo")
	fmt.Println(f.(func(string) string)("gophers"))
}
