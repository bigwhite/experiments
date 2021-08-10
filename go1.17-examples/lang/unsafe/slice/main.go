package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = [5]int{11, 12, 13, 14, 15}
	s1 := a[:]
	s2 := unsafe.Slice(&a[0], 5)

	fmt.Println(s1) // [11 12 13 14 15]
	fmt.Println(s2) // [11 12 13 14 15]
	fmt.Printf("the type of s2 is %T\n", s2)

	s2[2] += 10
	fmt.Println(a)  // [11 12 23 14 15]
	fmt.Println(s1) // [11 12 23 14 15]
	fmt.Println(s2) // [11 12 23 14 15]
}
