package main

import "./utils"

type I interface {
	IMethod1()
	IMethod2()
}

type T struct {
}

func (T) InstMethod() {

}
func (*T) PtrMethod() {

}

type MyInterface I
type MyStruct T

func main() {
	utils.DumpMethodSet((*I)(nil))

	var t T
	utils.DumpMethodSet(&t)
	var pt = &T{}
	utils.DumpMethodSet(&pt)

	utils.DumpMethodSet((*MyInterface)(nil))

	var m MyStruct
	utils.DumpMethodSet(&m)
	var pm = &MyStruct{}
	utils.DumpMethodSet(&pm)
}
