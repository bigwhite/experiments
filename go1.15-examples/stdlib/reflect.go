package main

import "reflect"

type u struct{}

func (u) M() { println("M") }

type t struct {
	u
	u2 u
}

func call(v reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	v.Method(0).Call(nil)
}

func main() {
	v := reflect.ValueOf(t{}) // v := t{}
	call(v)                   // v.M()
	call(v.Field(0))          // v.u.M()
	call(v.Field(1))          // v.u2.M()
}
