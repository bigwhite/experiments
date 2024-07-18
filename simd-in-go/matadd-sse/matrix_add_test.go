package main

import (
	"demo/pkg"
	"testing"
)

func TestMatrixAddNonSIMD(t *testing.T) {
	size := 1024
	a := make([]float32, size)
	b := make([]float32, size)
	c := make([]float32, size)
	expected := make([]float32, size)

	for i := 0; i < size; i++ {
		a[i] = float32(i)
		b[i] = float32(i)
		expected[i] = a[i] + b[i]
	}

	pkg.MatrixAddNonSIMD(a, b, c)

	for i := 0; i < size; i++ {
		if c[i] != expected[i] {
			t.Errorf("MatrixAddNonSIMD: expected %f, got %f at index %d", expected[i], c[i], i)
		}
	}
}

func TestMatrixAddSIMD(t *testing.T) {
	size := 1024
	a := make([]float32, size)
	b := make([]float32, size)
	c := make([]float32, size)
	expected := make([]float32, size)

	for i := 0; i < size; i++ {
		a[i] = float32(i)
		b[i] = float32(i)
		expected[i] = a[i] + b[i]
	}

	pkg.MatrixAddSIMD(a, b, c)

	for i := 0; i < size; i++ {
		if c[i] != expected[i] {
			t.Errorf("MatrixAddSIMD: expected %f, got %f at index %d", expected[i], c[i], i)
		}
	}
}
