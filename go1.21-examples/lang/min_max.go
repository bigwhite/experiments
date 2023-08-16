package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 5, 6
	fmt.Println(max(x))                    // 5
	fmt.Println(max(x, y, 0))              // 6
	fmt.Println(max("aby", "tony", "tom")) // tony

	var f float64 = 5.6
	// fmt.Printf("%T\n", max(x, y, f))    // invalid argument: mismatched types int (previous argument) and float64 (type of f)
	// fmt.Printf("%T\n", max(x, y, 10.1)) // (untyped float constant) truncated to int
	fmt.Println(max(f, math.NaN())) // NaN
	fmt.Println(min(f, math.NaN())) // NaN
}
