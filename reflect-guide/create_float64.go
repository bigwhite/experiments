package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := reflect.New(reflect.TypeOf(0.0))
	val.Elem().SetFloat(3.14)
	fmt.Println(val.Elem().Float()) // 输出：3.14
}
