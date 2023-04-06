package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {

	// define two two-rank tensor
	ta := tensor.New(tensor.WithBacking([]float32{1, 2, 3, 4}), tensor.WithShape(2, 2))
	fmt.Println("\ntensor a:")
	fmt.Println(ta)

	tb := tensor.New(tensor.WithBacking([]float32{5, 6, 7, 8}), tensor.WithShape(2, 2))
	fmt.Println("\ntensor b:")
	fmt.Println(tb)

	tc, err := tensor.Dot(ta, tb)
	if err != nil {
		fmt.Println("dot error:", err)
		return
	}
	fmt.Println("\ntensor a dot b:")
	fmt.Println(tc)

	td := tensor.New(tensor.WithBacking([]float32{5, 6, 7, 8, 9, 10}), tensor.WithShape(2, 3))
	fmt.Println("\ntensor d:")
	fmt.Println(td)
	te, err := tensor.Dot(ta, td)
	if err != nil {
		fmt.Println("dot error:", err)
		return
	}
	fmt.Println("\ntensor a dot d:")
	fmt.Println(te)

	// three-rank tensor dot two-rank tensor
	tf := tensor.New(tensor.WithBacking([]float32{23: 12}), tensor.WithShape(2, 3, 4))
	fmt.Println("\ntensor f:")
	fmt.Println(tf)

	tg := tensor.New(tensor.WithBacking([]float32{11: 12}), tensor.WithShape(4, 3))
	fmt.Println("\ntensor g:")
	fmt.Println(tg)

	th, err := tensor.Dot(tf, tg)
	if err != nil {
		fmt.Println("dot error:", err)
		return
	}
	fmt.Println("\ntensor f dot g:")
	fmt.Println(th)
}
