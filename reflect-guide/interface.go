package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func main() {
	var a Animal = Cat{}
	typ := reflect.TypeOf(a)
	fmt.Println(typ.Kind())         // interface
	fmt.Println(typ.NumMethod())    // 1
	fmt.Println(typ.Method(0).Name) // Speak
	fmt.Println(typ.Method(0).Type) // func(main.Animal) string

	fmt.Println(reflect.ValueOf(a).Type()) // main.Cat
}
