package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {

	// define two two-rank tensor
	ta := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2,
		2.7, 2.8, 1.5, 2.9,
		3.7, 2.4, 1.7, 3.1}), tensor.WithShape(3, 4))
	fmt.Println("\ntensor a:")
	fmt.Println(ta)

	tb := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2,
		2.7, 2.8, 1.5, 2.9,
		3.7, 2.4, 1.7, 3.1}), tensor.WithShape(3, 4))
	fmt.Println("\ntensor b:")
	fmt.Println(tb)

	tc, err := tensor.Mul(ta, tb)
	if err != nil {
		fmt.Println("multiply error:", err)
		return
	}
	fmt.Println("\ntensor a x b:")
	fmt.Println(tc)

	// multiple tensor and a scalar
	td, err := tensor.Mul(ta, float32(3.14))
	if err != nil {
		fmt.Println("multiply error:", err)
		return
	}
	fmt.Println("\ntensor ta x 3.14:")
	fmt.Println(td)

	td, err = tensor.Div(ta, tb)
	if err != nil {
		fmt.Println("divide error:", err)
		return
	}
	fmt.Println("\ntensor ta / tb:")
	fmt.Println(td)

	// multiply two tensors of different shape
	te := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3,
		3.2, 2.7, 2.8}), tensor.WithShape(2, 3))
	fmt.Println("\ntensor e:")
	fmt.Println(te)

	tf, err := tensor.Mul(ta, te)
	fmt.Println("\ntensor a x e:")
	if err != nil {
		fmt.Println("mul error:", err)
		return
	}
	fmt.Println(tf)
}
