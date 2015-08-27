package main

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d) //9,5,4,5
}
