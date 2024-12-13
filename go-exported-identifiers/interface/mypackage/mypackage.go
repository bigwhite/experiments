package mypackage

import "fmt"

type myStruct struct {
	Field string // 导出的字段
}

// NewMyStruct1是一个导出的函数，返回myStruct的指针
func NewMyStruct1(value string) *myStruct {
	return &myStruct{Field: value}
}

// NewMyStruct1是一个导出的函数，返回myStruct类型变量
func NewMyStruct2(value string) myStruct {
	return myStruct{Field: value}
}

func (m *myStruct) M1() {
	fmt.Println("invoke *myStruct's M1")
}

func (m myStruct) M2() {
	fmt.Println("invoke myStruct's M2")
}
