package main

/*
#include <stdio.h>
void plusOne(int **i) {
	(**i)++;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	sl := make([]*int, 5)
	var a int = 5
	sl[1] = &a
	C.plusOne((**C.int)((unsafe.Pointer)(&sl[0])))
	fmt.Println(sl[0])
}
