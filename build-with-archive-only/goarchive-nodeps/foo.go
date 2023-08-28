package foo

import "fmt"

func Add(a, b int) int {
	fmt.Println("invoke foo.Add")
	return a + b
}
