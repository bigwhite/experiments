package main

import (
	"fmt"
	_ "unsafe" // 必须导入 unsafe 包以使用 //go:linkname
)

// 声明符号链接
//
//go:linkname nanotime runtime.nanotime
func nanotime() int64

func main() {
	// 调用未导出的 runtime.nanotime 函数
	fmt.Println("Current time in nanoseconds:", nanotime())
}
