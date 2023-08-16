package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func negative[T Ordered](n T) bool {
	var zero T
	return n < zero
}

func main() {
	numbers := []int{0, 42, -10, 8}
	i := IndexFunc(numbers, negative)
	fmt.Println("First negative at index", i)
}
