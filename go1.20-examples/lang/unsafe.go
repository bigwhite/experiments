package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [6]byte{'h', 'e', 'l', 'l', 'o', '!'}
	s := unsafe.String(&arr[0], 6)
	fmt.Println(s) // hello!
	arr[0] = 'j'
	fmt.Println(s) // jello!

	b := unsafe.StringData(s)
	*b = 'k'
	fmt.Println(s) // kello!

	s1 := "golang"
	fmt.Println(s1) // golang
	b = unsafe.StringData(s1)
	*b = 'h' // fatal error: fault, unexpected fault address 0x10a67e5
	fmt.Println(s1)
}
