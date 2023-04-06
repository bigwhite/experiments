package main

import (
	"fmt"

	"gorgonia.org/tensor"
)

func main() {
	// define an one-rank tensor
	oneRankTensor := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2}), tensor.WithShape(4))
	fmt.Println("\none-rank tensor:")
	fmt.Println(oneRankTensor)
	fmt.Println("ndim:", oneRankTensor.Dims())
	fmt.Println("shape:", oneRankTensor.Shape())
	fmt.Println("dtype", oneRankTensor.Dtype())

	// define an two-rank tensor
	twoRankTensor := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2,
		2.7, 2.8, 1.5, 2.9,
		3.7, 2.4, 1.7, 3.1}), tensor.WithShape(3, 4))
	fmt.Println("\ntwo-rank tensor:")
	fmt.Println(twoRankTensor)
	fmt.Println("ndim:", twoRankTensor.Dims())
	fmt.Println("shape:", twoRankTensor.Shape())
	fmt.Println("dtype", twoRankTensor.Dtype())

	// define an three-rank tensor
	threeRankTensor := tensor.New(tensor.WithBacking([]float32{1.7, 2.6, 1.3, 3.2,
		2.7, 2.8, 1.5, 2.9,
		3.7, 2.4, 1.7, 3.1,
		1.5, 2.7, 1.4, 3.3,
		2.5, 2.8, 1.9, 2.9,
		3.5, 2.5, 1.7, 3.6}), tensor.WithShape(2, 3, 4))
	fmt.Println("\nthree-rank tensor:")
	fmt.Println(threeRankTensor)
	fmt.Println("ndim:", threeRankTensor.Dims())
	fmt.Println("shape:", threeRankTensor.Shape())
	fmt.Println("dtype", threeRankTensor.Dtype())

	// define an one-rank tensor which dtype is float64
	float64Tensor := tensor.New(tensor.WithBacking([]float64{1.7, 2.6, 1.3, 3.2}), tensor.WithShape(4))
	fmt.Println("\none-rank tensor with dtype float64:")
	fmt.Println(float64Tensor)
	fmt.Println("ndim:", float64Tensor.Dims())
	fmt.Println("shape:", float64Tensor.Shape())
	fmt.Println("dtype", float64Tensor.Dtype())
}
