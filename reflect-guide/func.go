package main

import (
	"fmt"
	"reflect"
)

func foo(a, b int, c *int) (int, bool) {
	*c = a + b
	return *c, true
}

func main() {
	typ := reflect.TypeOf(foo)
	fmt.Println(typ.Kind())                      // func
	fmt.Println(typ.NumIn())                     // 3
	fmt.Println(typ.In(0), typ.In(1), typ.In(2)) // int int *int
	fmt.Println(typ.NumOut())                    // 2
	fmt.Println(typ.Out(0))                      // int
	fmt.Println(typ.Out(1))                      // bool
}
