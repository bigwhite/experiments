// dot_simd.go
package main

import "simd"

// AVX2 版本，使用 256-bit 向量
func dotSIMD_AVX2(a, b []float32) float32 {
	const VEC_WIDTH = 8 // 使用 Float32x8
	var sumVec simd.Float32x8
	lenA := len(a)
	for i := 0; i <= lenA-VEC_WIDTH; i += VEC_WIDTH {
		va := simd.LoadFloat32x8Slice(a[i:])
		vb := simd.LoadFloat32x8Slice(b[i:])
		sumVec = sumVec.Add(va.Mul(vb))
	}
	var sumArr [VEC_WIDTH]float32
	sumVec.StoreSlice(sumArr[:])
	var sum float32
	for _, v := range sumArr {
		sum += v
	}
	for i := (lenA / VEC_WIDTH) * VEC_WIDTH; i < lenA; i++ {
		sum += a[i] * b[i]
	}
	return sum
}

// AVX 版本，使用 128-bit 向量
func dotSIMD_AVX(a, b []float32) float32 {
	const VEC_WIDTH = 4 // 使用 Float32x4
	var sumVec simd.Float32x4
	lenA := len(a)
	for i := 0; i <= lenA-VEC_WIDTH; i += VEC_WIDTH {
		va := simd.LoadFloat32x4Slice(a[i:])
		vb := simd.LoadFloat32x4Slice(b[i:])
		sumVec = sumVec.Add(va.Mul(vb))
	}
	var sumArr [VEC_WIDTH]float32
	sumVec.StoreSlice(sumArr[:])
	var sum float32
	for _, v := range sumArr {
		sum += v
	}
	for i := (lenA / VEC_WIDTH) * VEC_WIDTH; i < lenA; i++ {
		sum += a[i] * b[i]
	}
	return sum
}


// 调度函数
func dotSIMD(a, b []float32) float32 {
    if simd.HasAVX2() {
        return dotSIMD_AVX2(a, b)
    }
    // 注意：AVX是x86-64-v3的一部分，现代CPU普遍支持。
    // 为简单起见，这里假设AVX可用。生产代码中可能需要更细致的检测。
    return dotSIMD_AVX(a, b)
}
