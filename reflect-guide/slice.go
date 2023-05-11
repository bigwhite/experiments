package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]int, 5, 10)
	typ := reflect.TypeOf(s)
	fmt.Println(typ.Kind()) // slice
	fmt.Println(typ.Elem()) // int
	//fmt.Println(typ.Len())

	val := reflect.ValueOf(s)
	fmt.Println(val.Len()) // 5
	fmt.Println(val.Cap()) // 10
}
