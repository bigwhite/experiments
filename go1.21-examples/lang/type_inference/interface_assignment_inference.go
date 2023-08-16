package main

import "fmt"

type Indexable[T any] interface {
	At(i int) (T, bool)
}

func Index[T any](elems Indexable[T], i int) (T, bool) {
	return elems.At(i)
}

type MyList[T any] []T

func (m MyList[T]) At(i int) (T, bool) {
	var zero T
	if i > len(m) {
		return zero, false
	}
	return m[i], true
}

func main() {
	var m = MyList[int]{11, 12, 13}
	fmt.Println(Index(m, 2))
}
