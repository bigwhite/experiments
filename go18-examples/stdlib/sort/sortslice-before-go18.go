package main

import (
	"fmt"
	"sort"
)

type Lang struct {
	Name string
	Rank int
}

type TiboeIndexByRank []Lang

func (l TiboeIndexByRank) Len() int           { return len(l) }
func (l TiboeIndexByRank) Less(i, j int) bool { return l[i].Rank < l[j].Rank }
func (l TiboeIndexByRank) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func main() {
	langs := []Lang{
		{"rust", 2},
		{"go", 1},
		{"swift", 3},
	}
	sort.Sort(TiboeIndexByRank(langs))
	fmt.Printf("%v\n", langs)
}
