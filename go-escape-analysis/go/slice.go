package main

import (
	"reflect"
	"unsafe"
)

func noEscapeSliceWithDataInHeap() {
	var sl []int
	println("addr of local(noescape, data in heap) slice = ", &sl)
	printSliceHeader(&sl)
	sl = append(sl, 1)
	println("append 1")
	printSliceHeader(&sl)
	println("append 2")
	sl = append(sl, 2)
	printSliceHeader(&sl)
	println("append 3")
	sl = append(sl, 3)
	printSliceHeader(&sl)
	println("append 4")
	sl = append(sl, 4)
	printSliceHeader(&sl)
}

func noEscapeSliceWithDataInStack() {
	var sl = make([]int, 0, 8)
	println("addr of local(noescape, data in stack) slice = ", &sl)
	printSliceHeader(&sl)
	sl = append(sl, 1)
	println("append 1")
	printSliceHeader(&sl)
	sl = append(sl, 2)
	println("append 2")
	printSliceHeader(&sl)
}

func escapeSlice() *[]int {
	var sl = make([]int, 0, 8)
	println("addr of local(escape) slice = ", &sl)
	printSliceHeader(&sl)
	sl = append(sl, 1)
	println("append 1")
	printSliceHeader(&sl)
	sl = append(sl, 2)
	println("append 2")
	printSliceHeader(&sl)
	return &sl
}

func printSliceHeader(p *[]int) {
	ph := (*reflect.SliceHeader)(unsafe.Pointer(p))
	println("slice data =", unsafe.Pointer(ph.Data))
}

func main() {
	noEscapeSliceWithDataInHeap()
	noEscapeSliceWithDataInStack()
	escapeSlice()
}
