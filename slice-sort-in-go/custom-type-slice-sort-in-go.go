package main

import (
	"fmt"
	"sort"
)

type Lang struct {
	Name string
	Rank int
}

func main() {
	langs := []Lang{
		{"rust", 2},
		{"go", 1},
		{"swift", 3},
	}
	sort.Slice(langs, func(i, j int) bool { return langs[i].Rank < langs[j].Rank })
	fmt.Printf("%v\n", langs)
}
