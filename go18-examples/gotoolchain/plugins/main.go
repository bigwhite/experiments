package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, _ := plugin.Open("foo.so")
	f, _ := p.Lookup("Foo")
	fmt.Println(f.(func(string) string)("gophers"))
}
