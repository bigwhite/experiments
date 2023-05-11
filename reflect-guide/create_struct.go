package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", p.Name, p.Age)
}

func (p Person) SayHello(name string) {
	fmt.Printf("Hello, %s! My name is %s\n", name, p.Name)
}

func main() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(0),
		},
	})
	ptrVal := reflect.New(typ)
	val := ptrVal.Elem()
	val.FieldByName("Name").SetString("Alice")
	val.FieldByName("Age").SetInt(25)

	person := (*Person)(ptrVal.UnsafePointer())
	person.Greet()         // 输出：Hello, my name is Alice and I am 25 years old
	person.SayHello("Bob") // 输出：Hello, Bob! My name is Alice
}
