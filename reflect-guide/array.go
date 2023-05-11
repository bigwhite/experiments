package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	typ := reflect.TypeOf(arr)
	fmt.Println(typ.Kind())       // array
	fmt.Println(typ.Len())        // 5
	fmt.Println(typ.Comparable()) // true

	elemTyp := typ.Elem()
	fmt.Println(elemTyp.Kind())       // int
	fmt.Println(elemTyp.Comparable()) // true
}
