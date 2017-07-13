package main

import "fmt"

// type alias
type MyInt = int
type MyInt1 = MyInt

func main() {
	var i int = 5
	var mi MyInt = 6
	var mi1 MyInt1 = 7

	mi = i 
	mi1 = i
	mi1 = mi

	fmt.Println(i, mi, mi1)
}
