package main

import (
	"demo/mypackage"
)

// 定义一个导出的接口
type MyInterface interface {
	M1()
	M2()
}

func main() {
	var mi MyInterface

	// 通过导出的函数获取myStruct的指针
	ms1 := mypackage.NewMyStruct1("Hello1")
	mi = ms1
	mi.M1()
	mi.M2()

	// 通过导出的函数获取myStruct类型变量
	//ms2 := mypackage.NewMyStruct2("Hello2")
	//mi = ms2 // compile error: mypackage.myStruct does not implement MyInterface
}
