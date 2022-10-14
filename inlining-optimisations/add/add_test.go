package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = add(n, i)
	}
}
