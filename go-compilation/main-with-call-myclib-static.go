package main

/*
#cgo CFLAGS: -I ./my-c-lib
#cgo LDFLAGS: -static -L my-c-lib -lmylib
#include "mylib.h"
*/
import "C"
import "fmt"

func main() {
	// 调用 C 函数
	C.hello()

	// 调用 C 中的加法函数
	result := C.add(3, 4)
	fmt.Printf("Result of addition: %d\n", result)
}
