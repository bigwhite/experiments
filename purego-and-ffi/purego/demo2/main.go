package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/ebitengine/purego"
)

func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/lib/libSystem.B.dylib"
	case "linux":
		return "libc.so.6"
	// Windows 等其他平台...
	default:
		panic(fmt.Errorf("unsupported platform: %s", runtime.GOOS))
	}
}

func main() {
	libc, err := purego.Dlopen(getSystemLibrary(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	defer purego.Dlclose(libc)

	// 1. 定义与 C 函数 qsort 签名匹配的 Go 函数变量
	// void qsort(void *base, size_t nel, size_t width, int (*compar)(const void *, const void *));
	// 注意：最后一个参数应该是 uintptr，表示 C 函数指针
	var qsort func(data unsafe.Pointer, nitems uintptr, size uintptr, compar uintptr)
	purego.RegisterLibFunc(&qsort, libc, "qsort")

	// 2. 编写 Go 回调函数，签名必须与 qsort 的比较器兼容
	compareInts := func(a, b unsafe.Pointer) int {
		valA := *(*int)(a)
		valB := *(*int)(b)
		if valA < valB {
			return -1
		}
		if valA > valB {
			return 1
		}
		return 0
	}

	data := []int{88, 56, 100, 2, 25}
	fmt.Println("Original data:", data)

	// 3. 调用 qsort
	// 使用 NewCallback 将 Go 函数转换为 C 可调用的函数指针
	qsort(
		unsafe.Pointer(&data[0]),
		uintptr(len(data)),
		unsafe.Sizeof(int(0)),
		purego.NewCallback(compareInts),
	)

	fmt.Println("Sorted data:  ", data)

	// 验证结果
	if !reflect.DeepEqual(data, []int{2, 25, 56, 88, 100}) {
		panic("sort failed!")
	}
}
