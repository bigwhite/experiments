package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := reflect.New(reflect.TypeOf(""))
	val.Elem().SetString("hello")
	fmt.Println(val.Elem().String()) // 输出：hello
}
