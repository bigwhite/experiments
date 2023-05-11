package main

import (
	"fmt"
	"reflect"
)

func main() {
	typ := reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0))
	val := reflect.MakeChan(typ, 0)
	go func() {
		val.Send(reflect.ValueOf(42))
	}()

	ch, ok := val.Interface().(chan int)
	if !ok {
		fmt.Println("ch is not a chan int")
		return
	}
	fmt.Println(<-ch)
}
