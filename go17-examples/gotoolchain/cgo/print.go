package main

// #include <stdio.h>
// #include <stdlib.h>
//
// void print(void *array, int len) {
// 	char *c = (char*)array;
//
// 	for (int i = 0; i < len; i++) {
// 		printf("%c", *(c+i));
// 	}
// 	printf("\n");
// }
import "C"

import "unsafe"

func main() {
	var s = "hello cgo"
	csl := C.CBytes([]byte(s))
	C.print(csl, C.int(len(s)))
	C.free(unsafe.Pointer(csl))
}
