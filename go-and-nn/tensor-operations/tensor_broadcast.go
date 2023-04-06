package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {

	// broadcast a scalar
	ta := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2,
		2.7, 2.8, 1.5, 2.9,
		3.7, 2.4, 1.7, 3.1}), tensor.WithShape(3, 4))
	fmt.Println("\ntensor a:")
	fmt.Println(ta)

	tb, err := tensor.Mul(ta, float32(3.14))
	if err != nil {
		fmt.Println("broadcast error:", err)
		return
	}
	fmt.Println("\ntensor ta x 3.14:")
	fmt.Println(tb)

	tc := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2}), tensor.WithShape(2, 2))
	fmt.Println("\ntensor c:")
	fmt.Println(tc)

	td := tensor.New(tensor.WithBacking([]float32{10, 20}), tensor.WithShape(2))
	fmt.Println("\ntensor d:")
	fmt.Println(td)

	te, err := tensor.Mul(tc, td)
	if err != nil {
		fmt.Println("broadcast error:", err)
		return
	}
	fmt.Println("\ntensor tc x td:")
	fmt.Println(te)
}
