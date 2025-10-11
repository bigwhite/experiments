package main

import (
	"math/rand"
	"testing"
)

func BenchmarkPrefetching(b *testing.B) {
	size := 1024 * 1024
	data := make([]int, size)
	indices := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
		indices[i] = i
	}
	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	b.Run("Linear Access", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumLinear(data)
		}
	})

	b.Run("Random Access", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SumRandom(data, indices)
		}
	})
}
