package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("未获取到构建信息，请确保使用 Go Modules 构建")
		return
	}
	fmt.Printf("主模块: %s\n", info.Main.Path)
	fmt.Printf("Go版本: %s\n", info.GoVersion)
}
