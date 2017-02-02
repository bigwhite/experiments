package main

import (
	"fmt"
	"plugin"
)

func init() {
	fmt.Println("init in main program")
}

func main() {
	var err error
	fmt.Println("before opening the foo.so")

	p, err := plugin.Open("foo.so")
	if err != nil {
		fmt.Println("plugin Open error:", err)
		return
	}
	fmt.Println("after opening the foo.so")

	f, err := p.Lookup("Foo")
	if err != nil {
		fmt.Println("plugin Lookup symbol Foo error:", err)
	} else {
		fmt.Println(f.(func(string) string)("gophers"))
	}

	f, err = p.Lookup("foo")
	if err != nil {
		fmt.Println("plugin Lookup symbol foo error:", err)
	} else {
		fmt.Println(f.(func(string) string)("gophers"))
	}

	v, err := p.Lookup("V")
	if err != nil {
		fmt.Println("plugin Lookup symbol V error:", err)
	} else {
		fmt.Println(*v.(*int))
	}

	v, err = p.Lookup("v")
	if err != nil {
		fmt.Println("plugin Lookup symbol v error:", err)
	} else {
		fmt.Println(*v.(*int))
	}
}
