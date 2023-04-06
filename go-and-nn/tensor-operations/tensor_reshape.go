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
	fmt.Println("ta shape:", ta.Shape())

	err := ta.Reshape(2, 3)
	if err != nil {
		fmt.Println("reshape error:", err)
		return
	}
	fmt.Println("\ntensor a reshape(2,3):")
	fmt.Println(ta)
	fmt.Println("ta shape:", ta.Shape())

	err = ta.Reshape(1, 6)
	if err != nil {
		fmt.Println("reshape error:", err)
		return
	}
	fmt.Println("\ntensor a reshape(1, 6):")
	fmt.Println(ta)
	fmt.Println("ta shape:", ta.Shape())

	err = ta.Reshape(2, 1, 3)
	if err != nil {
		fmt.Println("reshape error:", err)
		return
	}
	fmt.Println("\ntensor a reshape(2, 1, 3):")
	fmt.Println(ta)
	fmt.Println("ta shape:", ta.Shape())
}
