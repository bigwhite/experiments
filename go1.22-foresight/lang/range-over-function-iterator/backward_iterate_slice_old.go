package main

import "fmt"

func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
		return
	}
}

func main() {
	sl := []string{"hello", "world", "golang"}
	Backward(sl)(func(i int, s string) bool {
		fmt.Printf("%d : %s\n", i, s)
		return true
	})
}
