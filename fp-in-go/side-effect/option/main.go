package main

import "fmt"

type Option[T any] struct {
	value   T
	present bool
}

func Some[T any](x T) Option[T] {
	return Option[T]{value: x, present: true}
}

func None[T any]() Option[T] {
	return Option[T]{present: false}
}

func (o Option[T]) Bind(f func(T) Option[T]) Option[T] {
	if !o.present {
		return None[T]()
	}
	return f(o.value)
}

// 使用示例
func safeDivide(a, b int) Option[int] {
	if b == 0 {
		return None[int]()
	}
	return Some(a / b)
}

func main() {
	result := Some(10).Bind(func(x int) Option[int] {
		return safeDivide(x, 2)
	})
	fmt.Println(result)

	result = Some(10).Bind(func(x int) Option[int] {
		return safeDivide(x, 0)
	})
	fmt.Println(result)
}
