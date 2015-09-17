package main

import "fmt"

func arrayRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("arrayRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println("")
}

func pointerToArrayRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("pointerToArrayRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println("")
}

func sliceRangeExpression() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("sliceRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a[:] {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[i] = v
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	fmt.Println("")
}

func sliceLenChangeRangeExpression() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	fmt.Println("sliceLenChangeRangeExpression result:")
	fmt.Println("a = ", a)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
}

func main() {
	arrayRangeExpression()
	pointerToArrayRangeExpression()
	sliceRangeExpression()
	sliceLenChangeRangeExpression()
}
