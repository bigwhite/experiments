package main

import (
	"fmt"
	"unsafe"
)

func slice2arrayptr() {
	var b = []int{11, 12, 13}
	var p = (*[3]int)(b)
	p[1] = p[1] + 10
	fmt.Printf("%v\n", b)
}
func slice2arrayptr1() {
	var b = []int{11, 12, 13}
	var p0 = (*[0]int)(b)
	fmt.Printf("%v\n", *p0)
	var p1 = (*[1]int)(b)
	fmt.Printf("%v\n", *p1)
	var p2 = (*[2]int)(b)
	fmt.Printf("%v\n", *p2)
}

func slice2arrayptrWithPanic() {
	var b = []int{11, 12, 13}
	//var p = (*[4]int)(b)
	var p = (*[3]int)(b[:1])
	fmt.Printf("%v\n", *p)
}

/*
// this can cause compile error
func slice2array() {
	var b = []int{11, 12, 13}
	var a = [3]int(b)
	fmt.Printf("%v\n", a)
}
*/

func slice2arrayptrWithHack() {
	var b = []int{11, 12, 13}
	var p = (*[3]int)(unsafe.Pointer(&b[0]))
	p[1] += 10
	fmt.Printf("%v\n", b) // [11 22 13]
}

func slice2arrayWithHack() {
	var b = []int{11, 12, 13}
	var a = *(*[3]int)(unsafe.Pointer(&b[0]))
	a[1] += 10
	fmt.Printf("%v\n", b) // [11 12 13]
}

func array2slice() {
	var a = [5]int{11, 12, 13, 14, 15}
	var b = a[0:len(a)] // or var b = a[:]
	b[1] += 10
	fmt.Printf("%v\n", b) // [11 22 13 14 15]
}

func main() {
	slice2arrayptrWithPanic()
}
