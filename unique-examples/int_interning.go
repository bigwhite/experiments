package main

import "unique"

func main() {
	var a, b int = 5, 6
	h1 := unique.Make(a)
	h2 := unique.Make(a)
	h3 := unique.Make(b)
	println(h1 == h2)
	println(h1 == h3)
}
