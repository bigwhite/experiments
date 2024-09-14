package main

import "unique"

func main() {
	h1 := unique.Make("hello")
	h2 := unique.Make("hello")
	h3 := unique.Make("hello")
	h4 := unique.Make("golang")
	println(h1 == h2)
	println(h1 == h3)
	println(h1 == h4)
	println(h2 == h4)
}
