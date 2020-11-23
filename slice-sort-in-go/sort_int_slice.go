package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	sl := IntSlice([]int{89, 14, 8, 9, 17, 56, 95, 3})
	fmt.Println(sl)
	sort.Sort(sl)
	fmt.Println(sl)
}
