package main

import (
	"fmt"
	"iter"
	"slices"
)

// Sequence 是一个包装 iter.Seq 的结构体，用于支持链式调用
type Sequence[T any] struct {
	seq iter.Seq[T]
}

// From 创建一个新的 Sequence
func From[T any](seq iter.Seq[T]) Sequence[T] {
	return Sequence[T]{seq: seq}
}

// Filter 方法
func (s Sequence[T]) Filter(f func(T) bool) Sequence[T] {
	return Sequence[T]{
		seq: func(yield func(T) bool) {
			for v := range s.seq {
				if f(v) && !yield(v) {
					return
				}
			}
		},
	}
}

// Map 方法
func (s Sequence[T]) Map(f func(T) T) Sequence[T] {
	return Sequence[T]{
		seq: func(yield func(T) bool) {
			for v := range s.seq {
				if !yield(f(v)) {
					return
				}
			}
		},
	}
}

// Range 方法，用于支持 range 语法
func (s Sequence[T]) Range() iter.Seq[T] {
	return s.seq
}

// 辅助函数
func IsEven(n int) bool {
	return n%2 == 0
}

func Add100(n int) int {
	return n + 100
}

func main() {
	sl := []int{12, 13, 14, 5, 67, 82}

	for v := range From(slices.Values(sl)).
		Filter(IsEven).
		Map(Add100).
		Range() {
		fmt.Println(v)
	}
}
