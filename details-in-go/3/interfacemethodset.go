package main

import (
	"fmt"
	"reflect"
)

type I interface {
	Method1()
	Method2()
}

func main() {
	var i *I
	elemType := reflect.TypeOf(i).Elem()
	n := elemType.NumMethod()
	for i := 0; i < n; i++ {
		fmt.Println(elemType.Method(i).Name)
	}
}
