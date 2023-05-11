package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.SliceOf(reflect.TypeOf(0))
	val := reflect.MakeSlice(typ, 3, 3)
	val.Index(0).SetInt(1)
	val.Index(1).SetInt(2)
	val.Index(2).SetInt(3)
	fmt.Println(val.Interface()) // 输出：[1 2 3]

	sl, ok := val.Interface().([]int)
	if !ok {
		fmt.Println("sl is not a []int")
		return
	}
	fmt.Println(sl) // [1 2 3]
}
