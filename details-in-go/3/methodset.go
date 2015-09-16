package main

import (
	"fmt"
	"reflect"
)

type S struct {
}

func (s S) InstMethodOfS_1() {
}

func (s S) InstMethodOfS_2() {
}

func (s *S) PtrMethodOfS_1() {
}

type T struct {
}

func (t T) T_Method1() {
}

func (t *T) T_Method2_PTR() {
}

type C struct {
	S
	*T
}

type I interface {
	S_Method1()
	S_Method2()
}

func dumpMethodSet(i interface{}) {
	fmt.Printf("=====%T 's Method Set: =======\n", i)

	v := reflect.TypeOf(i)
	n := v.NumMethod()

	for j := 0; j < n; j++ {
		fmt.Println(v.Method(j).Name)
	}

	fmt.Printf("=====%T 's Method Set end =======\n\n", i)
}

func main() {
	s := S{}
	dumpMethodSet(s)

	sPtr := &S{}
	dumpMethodSet(sPtr)

	t := T{}
	dumpMethodSet(t)

	tPtr := &T{}
	dumpMethodSet(tPtr)

	c := C{}
	dumpMethodSet(c)

	cPtr := &C{}
	dumpMethodSet(cPtr)

	var i I = s
	dumpMethodSet(i)

	i = sPtr
	dumpMethodSet(i)
}
