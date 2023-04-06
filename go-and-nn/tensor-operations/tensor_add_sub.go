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

	tc, _ := tensor.Add(ta, tb)
	fmt.Println("\ntensor a+b:")
	fmt.Println(tc)

	td, _ := tensor.Sub(ta, tb)
	fmt.Println("\ntensor a-b:")
	fmt.Println(td)

	// add in-place
	tensor.Add(ta, tb, tensor.UseUnsafe())
	fmt.Println("\ntensor a+b(in-place):")
	fmt.Println(ta)

	// tensor add scalar
	tg, err := tensor.Add(tb, float32(3.14))
	if err != nil {
		fmt.Println("add scalar error:", err)
		return
	}
	fmt.Println("\ntensor b+3.14:")
	fmt.Println(tg)

	// add two tensors of different shape
	te := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3,
		3.2, 2.7, 2.8}), tensor.WithShape(2, 3))
	fmt.Println("\ntensor e:")
	fmt.Println(te)

	tf, err := tensor.Add(ta, te)
	fmt.Println("\ntensor a+e:")
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	fmt.Println(tf)

}
