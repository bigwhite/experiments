// https://go.dev/play/p/ffxygzIdmCB?v=gotip
package main

import (
	"fmt"
	"slices"
)

type Seq0 func(yield func() bool)

func iter0[Slice ~[]E, E any](s Slice) Seq0 {
	return func(yield func() bool) {
		for range s {
			if !yield() {
				return
			}
		}
	}
}

var sl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {

	// 1. for range f {...}
	count := 0
	for range iter0(sl) {
		count++
	}
	fmt.Printf("total count = %d ", count)

	fmt.Printf("\n\n")

	// 2. for x := range f {...}
	fmt.Println("all values:")
	for v := range slices.Values(sl) {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n\n")

	// 3. for x,y := range f{...}
	fmt.Println("backward values:")
	for _, v := range slices.Backward(sl) {
		fmt.Printf("%d ", v)
	}
}
