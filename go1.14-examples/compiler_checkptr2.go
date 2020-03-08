package main

import (
	"unsafe"
)

func main() {
	var n = 5
	b := make([]byte, n)
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n+10))
	_ = end
}
