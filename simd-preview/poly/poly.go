package main

import "simd"

// Coefficients for our polynomial: y = 2.5x³ + 1.5x² + 0.5x + 3.0
const (
	c3 float32 = 2.5
	c2 float32 = 1.5
	c1 float32 = 0.5
	c0 float32 = 3.0
)

// polynomialScalar is the standard Go implementation, serving as our baseline.
// It uses Horner's method for efficient calculation.
func polynomialScalar(x []float32, y []float32) {
	for i, val := range x {
		res := (c3*val+c2)*val + c1
		y[i] = res*val + c0
	}
}

// polynomialSIMD_AVX uses 128-bit AVX instructions to process 4 floats at a time.
func polynomialSIMD_AVX(x []float32, y []float32) {
	const VEC_WIDTH = 4 // 128 bits / 32 bits per float = 4
	lenX := len(x)

	// Broadcast scalar coefficients to vector registers.
	// IMPORTANT: We manually create slices and use Load to avoid functions
	// like BroadcastFloat32x4 which might internally depend on AVX2.
	vc3 := simd.LoadFloat32x4Slice([]float32{c3, c3, c3, c3})
	vc2 := simd.LoadFloat32x4Slice([]float32{c2, c2, c2, c2})
	vc1 := simd.LoadFloat32x4Slice([]float32{c1, c1, c1, c1})
	vc0 := simd.LoadFloat32x4Slice([]float32{c0, c0, c0, c0})

	// Process the main part of the slice in chunks of 4.
	for i := 0; i <= lenX-VEC_WIDTH; i += VEC_WIDTH {
		vx := simd.LoadFloat32x4Slice(x[i:])

		// Apply Horner's method using SIMD vector operations.
		// vy = ((vc3 * vx + vc2) * vx + vc1) * vx + vc0
		vy := vc3.Mul(vx).Add(vc2)
		vy = vy.Mul(vx).Add(vc1)
		vy = vy.Mul(vx).Add(vc0)

		vy.StoreSlice(y[i:])
	}

	// Process any remaining elements at the end of the slice.
	for i := (lenX / VEC_WIDTH) * VEC_WIDTH; i < lenX; i++ {
		val := x[i]
		res := (c3*val+c2)*val + c1
		y[i] = res*val + c0
	}
}
