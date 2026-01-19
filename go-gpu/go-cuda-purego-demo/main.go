// main_pure.go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/ebitengine/purego"
)

const width = 1024

func main() {
	// 1. 加载动态库
	// 注意：在运行时，libmatrix.so 和 libcuder.so 必须在 LD_LIBRARY_PATH 中
	libMatrix, err := purego.Dlopen("libmatrix.so", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// 还需要加载 CUDA 运行时库，因为 libmatrix 依赖它
	_, err = purego.Dlopen("/usr/local/cuda/lib64/libcudart.so", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	// 2. 注册 C 函数符号
	var runMatrixMul func(a, b, c *float32, w int)
	purego.RegisterLibFunc(&runMatrixMul, libMatrix, "runMatrixMul")

	// 3. 准备数据 (与 CGO 版本相同)
	size := width * width
	h_a := make([]float32, size)
	h_b := make([]float32, size)
	h_c := make([]float32, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		h_a[i] = rand.Float32()
		h_b[i] = rand.Float32()
	}

	fmt.Println("Starting Matrix Multiplication via PureGo...")
	start := time.Now()

	// 4. 直接调用！无需 CGO 类型转换
	runMatrixMul(&h_a[0], &h_b[0], &h_c[0], width)

	// 5. 极其重要：保持内存存活
	// PureGo 调用是纯汇编实现，Go GC 无法感知堆栈上的指针引用
	// 必须显式保活，否则在计算期间 h_a 等可能被 GC 回收！
	runtime.KeepAlive(h_a)
	runtime.KeepAlive(h_b)
	runtime.KeepAlive(h_c)

	fmt.Printf("Done. Time: %v\n", time.Since(start))
	fmt.Printf("Result[0][0] = %f\n", h_c[0])
}
