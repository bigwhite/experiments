package main

import (
	"fmt"
	"reflect"
)

type X struct{}
type Y struct{}

func (*X) One()   { fmt.Println("hello 1") }
func (*X) Two()   { fmt.Println("hello 2") }
func (*X) Three() { fmt.Println("hello 3") }
func (*Y) Four()  { fmt.Println("hello 4") }
func (*Y) Five()  { fmt.Println("hello 5") }

func main() {
	var name string
	fmt.Scanf("%s", &name)
	reflect.ValueOf(&X{}).MethodByName(name).Call(nil)
	var y Y
	y.Five()
}
