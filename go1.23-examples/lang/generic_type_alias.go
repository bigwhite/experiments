package main

import "fmt"

type MySlice[T any] = []T

func main() {
	// 使用int类型实例化MySlice
	intSlice := MySlice[int]{1, 2, 3, 4, 5}
	fmt.Println("Int Slice:", intSlice)

	// 使用string类型实例化MySlice
	stringSlice := MySlice[string]{"hello", "world"}
	fmt.Println("String Slice:", stringSlice)

	// 使用自定义类型实例化MySlice
	type Person struct {
		Name string
		Age  int
	}

	personSlice := MySlice[Person]{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}

	fmt.Println("Person Slice:", personSlice)
}
