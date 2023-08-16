package main

import "fmt"

func main() {
	var sl = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("before clear, sl=%v, len(sl)=%d, cap(sl)=%d\n", sl, len(sl), cap(sl))
	clear(sl)
	fmt.Printf("after clear, sl=%v, len(sl)=%d, cap(sl)=%d\n", sl, len(sl), cap(sl))

	var m = map[string]int{
		"tony": 13,
		"tom":  14,
		"amy":  15,
	}
	fmt.Printf("before clear, m=%v, len(m)=%d\n", m, len(m))
	clear(m)
	fmt.Printf("after clear, m=%v, len(m)=%d\n", m, len(m))
}
