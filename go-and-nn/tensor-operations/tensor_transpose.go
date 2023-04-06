package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {

	// define two-rank tensor
	ta := tensor.New(tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6}), tensor.WithShape(3, 2))
	fmt.Println("\ntensor a:")
	fmt.Println(ta)

	tb, err := tensor.Transpose(ta)
	if err != nil {
		fmt.Println("transpose error:", err)
		return
	}
	fmt.Println("\ntensor a transpose:")
	fmt.Println(tb)

	// define three-rank tensor
	tc := tensor.New(tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6,
		7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18,
		19, 20, 21, 22, 23, 24}), tensor.WithShape(2, 3, 4))
	fmt.Println("\ntensor c:")
	fmt.Println(tc)
	fmt.Println("tc shape:", tc.Shape())

	td, err := tensor.Transpose(tc)
	if err != nil {
		fmt.Println("transpose error:", err)
		return
	}
	fmt.Println("\ntensor c transpose:")
	fmt.Println(td)
	fmt.Println("td shape:", td.Shape())
}
