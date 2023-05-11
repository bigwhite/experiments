package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	typ := reflect.PtrTo(reflect.TypeOf(Person{}))
	val := reflect.New(typ.Elem())
	val.Elem().FieldByName("Name").SetString("Alice")
	val.Elem().FieldByName("Age").SetInt(25)
	person := val.Interface().(*Person)
	fmt.Println(person.Name) // 输出：Alice
	fmt.Println(person.Age)  // 输出：25
}
