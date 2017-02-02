package main

import "fmt"

var V int
var v int

func init() {
	V = 17
	v = 23
	fmt.Println("init function in plugin foo")
}

func Foo(in string) string {
	return "Hello, " + in
}

func foo(in string) string {
	return "hello, " + in
}
