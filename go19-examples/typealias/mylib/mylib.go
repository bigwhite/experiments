package mylib

import "fmt"

type foo struct {
	A int
	B string
}

type Foo = foo

func (f *foo) String() {
	fmt.Println(f.A, f.B)
}

func (f *foo) anotherMethod() {

}
