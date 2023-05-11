package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := reflect.New(reflect.TypeOf(0))
	val.Elem().SetInt(42)
	fmt.Println(val.Elem().Int()) // 输出：42
}
