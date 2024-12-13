package mypackage

import "fmt"

type nonExported struct {
	Field string // 导出的字段
}

// Exported 是导出的结构体，嵌入了nonExported
type Exported struct {
	nonExported // 嵌入非导出结构体
}

func NewExported(value string) *Exported {
	return &Exported{
		nonExported: nonExported{
			Field: value,
		},
	}
}

// M1是导出的函数
func (n *nonExported) M1() {
	fmt.Println("invoke nonExported's M1")
}

// M2是导出的函数
func (e *Exported) M2() {
	fmt.Println("invoke Exported's M2")
}
