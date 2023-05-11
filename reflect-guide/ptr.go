package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	p := &i
	typ := reflect.TypeOf(p)
	fmt.Println(typ.Kind()) // ptr
	fmt.Println(typ.Elem()) // int

	fmt.Println(reflect.ValueOf(p).Elem().Int()) // 10
}
