package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	sex  string
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s, and I'm %d years old.\n", p.Name, p.Age)
}
func (p Person) unexportedMethod() {
}

func main() {
	p := Person{Name: "Tom", Age: 20, sex: "male"}
	typ := reflect.TypeOf(p)
	fmt.Println(typ.Kind())                   // struct
	fmt.Println(typ.NumField())               // 2
	fmt.Println(typ.Field(0).Name)            // Name
	fmt.Println(typ.Field(0).Type)            // string
	fmt.Println(typ.Field(0).Tag)             // json:"name"
	fmt.Println(typ.Field(1).Name)            // Age
	fmt.Println(typ.Field(1).Type)            // int
	fmt.Println(typ.Field(1).Tag)             // json:"age"
	fmt.Println(typ.Field(2).Name)            // sex
	fmt.Println(typ.Method(0).Name)           // SayHello
	fmt.Println(typ.Method(0).Type)           // func(main.Person)
	fmt.Println(typ.Method(0).Func)           // 0x109b6e0
	fmt.Println(typ.MethodByName("SayHello")) // {SayHello func(main.Person)}
	fmt.Println(typ.MethodByName("unexportedMethod"))
}
