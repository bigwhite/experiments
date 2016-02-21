package main

/*
#include <stdio.h>
struct Foo{
	int a;
	int *p;
};

void plusOne(struct Foo *f) {
	(f->a)++;
	*(f->p)++;
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
	//f.p = &f.a

	C.plusOne(f)
	fmt.Println(int(f.a))
}
