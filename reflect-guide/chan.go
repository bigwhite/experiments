package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan<- int, 10)
	ch <- 1
	ch <- 2
	typ := reflect.TypeOf(ch)
	fmt.Println(typ.Kind())    // chan
	fmt.Println(typ.Elem())    // int
	fmt.Println(typ.ChanDir()) // chan<-

	fmt.Println(reflect.ValueOf(ch).Len()) // 2
	fmt.Println(reflect.ValueOf(ch).Cap()) // 10
}
