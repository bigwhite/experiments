package mypackage

import "fmt"

// 定义一个非导出的结构体
type nonExported struct {
	Field string
}

// 导出的方法
func (n *nonExported) M1() {
	fmt.Println("invoke nonExported's M1")
}

func (n *nonExported) M2() {
	fmt.Println("invoke nonExported's M2")
}

// 导出的函数，用于创建非导出类型的实例
func NewNonExported(value string) *nonExported {
	return &nonExported{Field: value}
}
