package main

import "fmt"

func init() {
	fmt.Println("init function in plugin foo")
}

func Foo(in string) string {
	return "hello, " + in
}
