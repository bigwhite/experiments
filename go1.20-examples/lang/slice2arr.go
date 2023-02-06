package main

import "fmt"

func slice2arrOK() {
	var sl = []int{1, 2, 3, 4, 5, 6, 7}
	var arr = [7]int(sl)
	var parr = (*[7]int)(sl)
	fmt.Println(sl)  // [1 2 3 4 5 6 7]
	fmt.Println(arr) // [1 2 3 4 5 6 7]
	sl[0] = 11
	fmt.Println(arr)  // [1 2 3 4 5 6 7]
	fmt.Println(parr) // &[11 2 3 4 5 6 7]
}

func slice2arrPanic() {
	var sl = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(sl)
	var arr = [8]int(sl) // panic: runtime error: cannot convert slice with length 7 to array or pointer to array with length 8
	fmt.Println(arr)     // &[11 2 3 4 5 6 7]

}

func main() {
	slice2arrOK()
	slice2arrPanic()
}
