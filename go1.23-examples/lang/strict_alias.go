package main

import "fmt"

type MySlice[T any] = []T
type YourSlice[T comparable] = MySlice[T]

func main() {
	// 使用int类型实例化MySlice
	intSlice := MySlice[int]{1, 2, 3, 4, 5}
	fmt.Println("Int Slice:", intSlice)

	intsliceSlice := YourSlice[[]int]{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println("IntSlice Slice:", intsliceSlice)
}
