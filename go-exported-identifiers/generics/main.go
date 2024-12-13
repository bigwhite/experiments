package main

import (
	"demo/mypackage"
)

// 定义一个用作约束的接口
type MyInterface interface {
	M1()
	M2()
}

func UseNonExportedAsTypeArgument[T MyInterface](item T) {
	item.M1()
	item.M2()
}

// 定义一个带有泛型参数的新类型
type GenericType[T MyInterface] struct {
	Item T
}

func NewGenericType[T MyInterface](item T) GenericType[T] {
	return GenericType[T]{Item: item}
}

func main() {
	// 创建非导出类型的实例
	n := mypackage.NewNonExported("Hello")

	// 调用泛型函数，传入实现了MyInterface的非导出类型
	UseNonExportedAsTypeArgument(n) // ok

	// g := GenericType{Item: n} // compiler error: cannot use generic type GenericType[T MyInterface] without instantiation
	g := NewGenericType(n)
	g.Item.M1()
}
