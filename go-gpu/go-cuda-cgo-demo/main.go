// main.go
package main

/*
#cgo LDFLAGS: -L. -lmatrix -L/usr/local/cuda/lib64 -lcudart
#include "matrix.h"
*/
import "C"
import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const width = 1024 // 矩阵大小 1024x1024，共 100万次计算

func main() {
	size := width * width
	h_a := make([]float32, size)
	h_b := make([]float32, size)
	h_c := make([]float32, size)

	// 初始化矩阵数据
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		h_a[i] = rand.Float32()
		h_b[i] = rand.Float32()
	}

	fmt.Printf("Starting Matrix Multiplication (%dx%d) on GPU...\n", width, width)
	start := time.Now()

	// 调用 CUDA 函数
	// 使用 unsafe.Pointer 获取切片的底层数组指针，传递给 C
	C.runMatrixMul(
		(*C.float)(unsafe.Pointer(&h_a[0])),
		(*C.float)(unsafe.Pointer(&h_b[0])),
		(*C.float)(unsafe.Pointer(&h_c[0])),
		C.int(width),
	)

	// 注意：在更复杂的场景中，需要使用 runtime.KeepAlive(h_a)
	// 来确保 Go GC 不会在 CGO 调用期间回收切片内存。

	elapsed := time.Since(start)
	fmt.Printf("Done. Time elapsed: %v\n", elapsed)

	// 简单验证：检查左上角元素
	fmt.Printf("Result[0][0] = %f\n", h_c[0])
}
