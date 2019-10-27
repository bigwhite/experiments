package main

import "fmt"

func main() {
	var i int = 5
	fmt.Println(2 << uint(i)) // before go 1.13
	fmt.Println(2 << i)       // in go 1.13 and later version
}
