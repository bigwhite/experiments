// dot_simd.go
package main

import "simd"

const VEC_WIDTH = 8 // 使用 AVX2 的 Float32x8，一次处理 8 个 float32

func dotSIMD(a, b []float32) float32 {
	var sumVec simd.Float32x8 // 累加和向量，初始为全 0
	lenA := len(a)

	// 处理能被 VEC_WIDTH 整除的主要部分
	for i := 0; i <= lenA-VEC_WIDTH; i += VEC_WIDTH {
		va := simd.LoadFloat32x8Slice(a[i:])
		vb := simd.LoadFloat32x8Slice(b[i:])
		
		// 向量乘法，然后累加到 sumVec
		sumVec = sumVec.Add(va.Mul(vb))
	}

	// 将累加和向量中的所有元素水平相加
	var sumArr [VEC_WIDTH]float32
	sumVec.StoreSlice(sumArr[:])
	var sum float32
	for _, v := range sumArr {
		sum += v
	}

	// 处理剩余的尾部元素
	for i := (lenA / VEC_WIDTH) * VEC_WIDTH; i < lenA; i++ {
		sum += a[i] * b[i]
	}

	return sum
}
