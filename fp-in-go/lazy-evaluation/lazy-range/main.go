package main

import "fmt"

func lazyRange(start, end int) func() (int, bool) {
	current := start
	return func() (int, bool) {
		if current >= end {
			return 0, false
		}
		result := current
		current++
		return result, true
	}
}
func main() {
	next := lazyRange(1, 5)
	for {
		value, hasNext := next()
		if !hasNext {
			break
		}
		fmt.Println(value)
	}
}
