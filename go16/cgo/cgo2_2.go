package main

// #include <stdlib.h>
// extern void goFoo(int**);
//
// static void cFoo() {
//     int **p = malloc(sizeof(int*));
//     goFoo(p);
// }
import "C"

//export goFoo
func goFoo(p **C.int) {
	*p = new(C.int)
}

func main() {
	C.cFoo()
}
