// main.go
package main

import (
	"demo/mypackage"
	"fmt"
)

func main() {
	// 通过导出的函数获取myStruct的指针
	ms1 := mypackage.NewMyStruct1("Hello1")

	// 尝试访问Field字段
	fmt.Println(ms1.Field) // ok

	// 通过导出的函数获取myStruct类型变量
	ms2 := mypackage.NewMyStruct1("Hello2")

	// 尝试访问Field字段
	fmt.Println(ms2.Field) // ok
}
