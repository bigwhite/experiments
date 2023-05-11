package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	val := reflect.MakeMap(typ)
	key1 := reflect.ValueOf("one")
	value1 := reflect.ValueOf(1)
	key2 := reflect.ValueOf("two")
	value2 := reflect.ValueOf(2)
	val.SetMapIndex(key1, value1)
	val.SetMapIndex(key2, value2)
	fmt.Println(val.Interface()) // 输出：map[one:1 two:2]

	m, ok := val.Interface().(map[string]int)
	if !ok {
		fmt.Println("m is not a map[string]int")
		return
	}

	fmt.Println(m)

}
