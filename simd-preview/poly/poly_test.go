package main

import (
	"math"
	"math/rand"
	"testing"
)

const sliceSize = 8192

var (
	sliceX []float32
	sliceY []float32 // A slice to write results into
)

func init() {
	sliceX = make([]float32, sliceSize)
	sliceY = make([]float32, sliceSize)
	for i := 0; i < sliceSize; i++ {
		sliceX[i] = rand.Float32() * 2.0 // Random floats between 0.0 and 2.0
	}
}

// checkFloats compares two float slices for near-equality.
func checkFloats(t *testing.T, got, want []float32, tolerance float64) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("slices have different lengths: got %d, want %d", len(got), len(want))
	}
	for i := range got {
		if math.Abs(float64(got[i]-want[i])) > tolerance {
			t.Errorf("mismatch at index %d: got %f, want %f", i, got[i], want[i])
			return
		}
	}
}

// TestPolynomialCorrectness ensures the SIMD implementation matches the scalar one.
func TestPolynomialCorrectness(t *testing.T) {
	yScalar := make([]float32, sliceSize)
	ySIMD := make([]float32, sliceSize)

	polynomialScalar(sliceX, yScalar)
	polynomialSIMD_AVX(sliceX, ySIMD)

	// Use a small tolerance for floating point comparisons.
	checkFloats(t, ySIMD, yScalar, 1e-6)
}

func BenchmarkPolynomialScalar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		polynomialScalar(sliceX, sliceY)
	}
}

func BenchmarkPolynomialSIMD_AVX(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		polynomialSIMD_AVX(sliceX, sliceY)
	}
}
