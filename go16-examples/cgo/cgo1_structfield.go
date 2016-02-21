package main

/*
#include <stdio.h>
struct Foo{
	int a;
	int *p;
};

void plusOne(int *i) {
	(*i)++;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	f := &C.struct_Foo{}
	f.a = 5
	f.p = (*C.int)((unsafe.Pointer)(new(int)))

	C.plusOne(&f.a)
	fmt.Println(int(f.a))
}
