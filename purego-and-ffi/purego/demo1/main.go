package main

import (
	"fmt"
	"runtime"

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
	// 1. 加载 C 库
	libc, err := purego.Dlopen(getSystemLibrary(), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	defer purego.Dlclose(libc) // 确保库被卸载

	// 2. 声明一个 Go 函数变量，其签名与 C 函数匹配
	var puts func(string)

	// 3. 注册！将 Go 变量与 C 函数 "puts" 绑定
	purego.RegisterLibFunc(&puts, libc, "puts")

	// 4. 直接像调用普通 Go 函数一样调用它！
	puts("Calling C from Go without CGO!")
}
