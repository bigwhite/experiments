package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	typ := reflect.TypeOf(m)
	fmt.Println(typ.Kind()) // map
	fmt.Println(typ.Key())  // string
	fmt.Println(typ.Elem()) // int

	fmt.Println(reflect.ValueOf(m).Len()) // 3
}
