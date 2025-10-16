package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/ffi"
)

// getSystemLibrary 函数与前一个示例相同
func getSystemLibrary() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/lib/libSystem.B.dylib"
	case "linux":
		return "libc.so.6"
	default:
		panic(fmt.Errorf("unsupported platform: %s", runtime.GOOS))
	}
}

// C 语言中的 struct timeval
//
//	struct timeval {
//	    time_t      tv_sec;     /* seconds */
//	    suseconds_t tv_usec;    /* microseconds */
//	};
//
// Go 版本的结构体，注意字段类型和大小必须与 C 版本兼容
// 在 64 位系统上，time_t 和 suseconds_t 通常都是 int64
type Timeval struct {
	TvSec  int64 // 秒
	TvUsec int64 // 微秒
}

func main() {
	libc, err := purego.Dlopen(getSystemLibrary(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	defer purego.Dlclose(libc)

	// 1. 获取 C 函数地址
	gettimeofday_addr, err := purego.Dlsym(libc, "gettimeofday")
	if err != nil {
		panic(err)
	}

	// 2. 使用 ffi.PrepCif 准备函数签名
	// int gettimeofday(struct timeval *tv, struct timezone *tz);
	// 返回值: int (ffi.TypeSint32)
	// 参数1: struct timeval* (ffi.TypePointer)
	// 参数2: struct timezone* (ffi.TypePointer)，我们传入 nil
	var cif ffi.Cif
	if status := ffi.PrepCif(&cif, ffi.DefaultAbi, 2, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypePointer); status != ffi.OK {
		panic(fmt.Sprintf("PrepCif failed with status: %v", status))
	}

	// 3. 准备 Go 结构体实例，用于接收 C 函数的输出
	var tv Timeval

	// 4. 准备参数
	// ffi.Call 需要一个指向参数的指针数组
	// 第一个参数：指向 Timeval 结构体的指针
	// 第二个参数：nil（表示 timezone 参数为 NULL）
	arg1 := unsafe.Pointer(&tv)
	var arg2 unsafe.Pointer = nil

	// 创建参数指针数组
	args := []unsafe.Pointer{
		unsafe.Pointer(&arg1),
		unsafe.Pointer(&arg2),
	}

	// 5. 调用 C 函数
	var ret int32
	ffi.Call(&cif, gettimeofday_addr, unsafe.Pointer(&ret), args...)

	if ret != 0 {
		panic(fmt.Sprintf("gettimeofday failed with return code: %d", ret))
	}

	// 6. 解释结果
	fmt.Printf("C gettimeofday result:\n")
	fmt.Printf("  - Seconds: %d\n", tv.TvSec)
	fmt.Printf("  - Microseconds: %d\n", tv.TvUsec)

	// 与 Go 标准库的结果进行对比
	goTime := time.Now()
	fmt.Printf("\nGo time.Now() result:\n")
	fmt.Printf("  - Seconds: %d\n", goTime.Unix())
	fmt.Printf("  - Microseconds component: %d\n", goTime.Nanosecond()/1000)

	// 验证秒数是否大致相等
	timeDiff := goTime.Unix() - tv.TvSec
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}
	if timeDiff > 1 {
		panic(fmt.Sprintf("seconds mismatch! Diff: %d", timeDiff))
	}
	fmt.Println("\nSuccess! The results are consistent.")
}
