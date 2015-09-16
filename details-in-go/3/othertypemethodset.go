package main

import "./utils"

type T struct {
}

func (t T) Method1() {
}

func (t *T) Method2() {
}

func (t *T) Method3() {
}

type I interface {
	Method1()
	Method2()
}

func main() {
	var t T
	utils.DumpMethodSet(&t)

	var pt = &T{}
	utils.DumpMethodSet(&pt)

	utils.DumpMethodSet((*I)(nil))

	// var i I = t // cannot use t (type T) as type I in assignment: T does not implement I (Method2 method has pointer receiver)
	var i I = pt
	i.Method1()

	pt.Method1()
	t.Method3()
}
