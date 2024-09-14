package main

import "unique"

type UserType struct {
	a int
	z float64
	s string
}

func main() {
	var u1 = UserType{
		a: 5,
		z: 3.14,
		s: "golang",
	}
	var u2 = UserType{
		a: 5,
		z: 3.15,
		s: "golang",
	}
	h1 := unique.Make(u1)
	h2 := unique.Make(u1)
	h3 := unique.Make(u2)
	println(h1 == h2)
	println(h1 == h3)
}
