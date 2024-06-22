package main

import (
	"iter"
	"slices"
)

// Filter returns an iterator over seq that only includes
// the values v for which f(v) is true.
func Filter[V any](f func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// 过滤奇数
func FilterOdd(seq iter.Seq[int]) iter.Seq[int] {
	return Filter[int](func(n int) bool {
		return n%2 == 0
	}, seq)
}

// Map returns an iterator over f applied to seq.
func Map[In, Out any](f func(In) Out, seq iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
	}
}

// Add 100 to every element in seq
func Add100(seq iter.Seq[int]) iter.Seq[int] {
	return Map[int, int](func(n int) int {
		return n + 100
	}, seq)
}

var sl = []int{12, 13, 14, 5, 67, 82}

func main() {
	for v := range Add100(FilterOdd(slices.Values(sl))) {
		println(v)
	}
}
