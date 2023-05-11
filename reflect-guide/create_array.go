package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.ArrayOf(3, reflect.TypeOf(0))
	val := reflect.New(typ)
	arr := val.Elem()
	arr.Index(0).SetInt(1)
	arr.Index(1).SetInt(2)
	arr.Index(2).SetInt(3)
	fmt.Println(arr.Interface()) // 输出：[1 2 3]
	arr1, ok := arr.Interface().([3]int)
	if !ok {
		fmt.Println("not a [3]int")
		return
	}

	fmt.Println(arr1)
}
