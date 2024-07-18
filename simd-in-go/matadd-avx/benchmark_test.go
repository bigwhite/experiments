package main

import (
	"demo/pkg"
	"testing"
)

func BenchmarkMatrixAddNonSIMD(tb *testing.B) {
	size := 1024
	a := make([]float32, size)
	b := make([]float32, size)
	c := make([]float32, size)

	for i := 0; i < size; i++ {
		a[i] = float32(i)
		b[i] = float32(i)
	}

	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		pkg.MatrixAddNonSIMD(a, b, c)
	}
}

func BenchmarkMatrixAddSIMD(tb *testing.B) {
	size := 1024
	a := make([]float32, size)
	b := make([]float32, size)
	c := make([]float32, size)

	for i := 0; i < size; i++ {
		a[i] = float32(i)
		b[i] = float32(i)
	}

	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		pkg.MatrixAddSIMD(a, b, c)
	}
}
