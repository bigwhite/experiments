package main

// extern int* goAdd(int, int);
//
// static int cAdd(int a, int b) {
//     int *i = goAdd(a, b);
//     return *i;
// }
import "C"
import "fmt"

//export goAdd
func goAdd(a, b C.int) *C.int {
	c := a + b
	return &c
}

func main() {
	var a, b int = 5, 6
	i := C.cAdd(C.int(a), C.int(b))
	fmt.Println(int(i))
}
