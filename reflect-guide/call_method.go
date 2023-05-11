package main

import (
	"fmt"
	"reflect"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area(factor float64) float64 {
	return r.Width * r.Height * factor
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	val := reflect.ValueOf(r)
	method := val.MethodByName("Area")
	args := []reflect.Value{reflect.ValueOf(1.5)}
	result := method.Call(args)
	fmt.Println(result[0].Float()) // 输出：75
}
