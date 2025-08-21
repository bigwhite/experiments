// dot_test.go
package main

import (
	"math/rand"
	"testing"
)

func generateSlice(n int) []float32 {
	s := make([]float32, n)
	for i := range s {
		s[i] = rand.Float32()
	}
	return s
}

var (
	sliceA = generateSlice(4096)
	sliceB = generateSlice(4096)
)

func BenchmarkDotScalar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dotScalar(sliceA, sliceB)
	}
}

func BenchmarkDotSIMD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dotSIMD(sliceA, sliceB)
	}
}
