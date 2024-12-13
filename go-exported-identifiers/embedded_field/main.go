package main

import (
	"demo/mypackage"
	"fmt"
)

// 定义一个导出的接口
type MyInterface interface {
	M1()
	M2()
}

func main() {
	ms := mypackage.NewExported("Hello")
	fmt.Println(ms.Field) // 访问嵌入的非导出结构体的导出字段

	ms.M1() // 访问嵌入的非导出结构体的导出方法

	var mi MyInterface = ms
	mi.M1()
	mi.M2()
}
