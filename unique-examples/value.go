package main

import (
	"fmt"
	"unique"
)

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
	h1 := unique.Make(u1)
	h2 := unique.Make("hello, golang")
	h3 := unique.Make(567890)
	v1 := h1.Value()
	v2 := h2.Value()
	v3 := h3.Value()
	fmt.Printf("%T: %v\n", v1, v1)
	fmt.Printf("%T: %v\n", v2, v2)
	fmt.Printf("%T: %v\n", v3, v3)
}
