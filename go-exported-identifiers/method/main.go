// main.go
package main

import (
	"demo/mypackage"
)

func main() {
	// 通过导出的函数获取myStruct的指针
	ms1 := mypackage.NewMyStruct1("Hello1")
	ms1.M1()
	ms1.M2()

	// 通过导出的函数获取myStruct类型变量
	ms2 := mypackage.NewMyStruct2("Hello2")
	ms2.M1()
	ms2.M2()
}
