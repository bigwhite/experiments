package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := []int{89, 14, 8, 9, 17, 56, 95, 3}
	fmt.Println(sl)
	sort.Ints(sl)
	fmt.Println(sl)
}
