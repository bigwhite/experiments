package main

import "fmt"

// Collection 接口定义了通用的集合操作
type Collection[T any] interface {
	Filter(predicate func(T) bool) Collection[T]
	Map(transform func(T) T) Collection[T]
	Reduce(initialValue T, reducer func(T, T) T) T
}

// SliceCollection 是基于切片的集合实现
type SliceCollection[T any] struct {
	data []T
}

// NewSliceCollection 创建一个新的 SliceCollection
func NewSliceCollection[T any](data []T) *SliceCollection[T] {
	return &SliceCollection[T]{data: data}
}

// Filter 实现了 Collection 接口的 Filter 方法
func (sc *SliceCollection[T]) Filter(predicate func(T) bool) Collection[T] {
	result := make([]T, 0)
	for _, item := range sc.data {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return &SliceCollection[T]{data: result}
}

// Map 实现了 Collection 接口的 Map 方法
func (sc *SliceCollection[T]) Map(transform func(T) T) Collection[T] {
	result := make([]T, len(sc.data))
	for i, item := range sc.data {
		result[i] = transform(item)
	}
	return &SliceCollection[T]{data: result}
}

// Reduce 实现了 Collection 接口的 Reduce 方法
func (sc *SliceCollection[T]) Reduce(initialValue T, reducer func(T, T) T) T {
	result := initialValue
	for _, item := range sc.data {
		result = reducer(result, item)
	}
	return result
}

// SetCollection 是基于 map 的集合实现
type SetCollection[T comparable] struct {
	data map[T]struct{}
}

// NewSetCollection 创建一个新的 SetCollection
func NewSetCollection[T comparable]() *SetCollection[T] {
	return &SetCollection[T]{data: make(map[T]struct{})}
}

// Add 向 SetCollection 添加元素
func (sc *SetCollection[T]) Add(item T) {
	sc.data[item] = struct{}{}
}

// Filter 实现了 Collection 接口的 Filter 方法
func (sc *SetCollection[T]) Filter(predicate func(T) bool) Collection[T] {
	result := NewSetCollection[T]()
	for item := range sc.data {
		if predicate(item) {
			result.Add(item)
		}
	}
	return result
}

// Map 实现了 Collection 接口的 Map 方法
func (sc *SetCollection[T]) Map(transform func(T) T) Collection[T] {
	result := NewSetCollection[T]()
	for item := range sc.data {
		result.Add(transform(item))
	}
	return result
}

// Reduce 实现了 Collection 接口的 Reduce 方法
func (sc *SetCollection[T]) Reduce(initialValue T, reducer func(T, T) T) T {
	result := initialValue
	for item := range sc.data {
		result = reducer(result, item)
	}
	return result
}

// ToSlice 实现了 Collection 接口的 ToSlice 方法
func (sc *SetCollection[T]) ToSlice() []T {
	result := make([]T, 0, len(sc.data))
	for item := range sc.data {
		result = append(result, item)
	}
	return result
}

func main() {
	// 使用 SliceCollection
	numbers := NewSliceCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	result := numbers.
		Filter(func(n int) bool { return n%2 == 0 }).
		Map(func(n int) int { return n * 2 }).
		Reduce(0, func(acc, n int) int { return acc + n })
	fmt.Println(result) // 输出: 60

	// 使用 SetCollection
	set := NewSetCollection[int]()
	for _, n := range []int{1, 2, 2, 3, 3, 3, 4, 5} {
		set.Add(n)
	}
	uniqueSum := set.
		Filter(func(n int) bool { return n > 2 }).
		Map(func(n int) int { return n * n }).
		Reduce(0, func(acc, n int) int { return acc + n })
	fmt.Println(uniqueSum) // 输出: 50 (3^2 + 4^2 + 5^2)
}
