package main

import (
	"fmt"
	"unsafe"
)

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func foo() {
	var a int = 66666666
	var b int = 77
	fmt.Printf("addr of a in bar = %p\n", (*int)(noescape(unsafe.Pointer(&a))))
	println("addr of a in bar =", &a)
	println("addr of b in bar =", &b)
}

func main() {
	foo()
}
