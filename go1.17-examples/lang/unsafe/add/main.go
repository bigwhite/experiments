package main

import (
	"fmt"
	"unsafe"
)

const intLen = unsafe.Sizeof(int(8))

func foo() {
	var a = [5]int{11, 12, 13, 14, 15}
	for i := 0; i < 5; i++ {
		p := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + uintptr(uintptr(i)*intLen)))
		*p = *p + 10
	}
	fmt.Println(a)
}

func bar() {
	var a = [5]int{11, 12, 13, 14, 15}
	for i := 0; i < 5; i++ {
		p := (*int)(unsafe.Add(unsafe.Pointer(&a[0]), uintptr(i)*intLen))
		*p = *p + 10
	}
	fmt.Println(a)
}

func main() {
	foo()
	bar()
}
