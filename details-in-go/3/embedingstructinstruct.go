package main

import "./utils"

type T struct {
}

func (T) InstMethod1OfT() {

}

func (T) InstMethod2OfT() {

}

func (*T) PtrMethodOfT() {

}

type S struct {
}

func (S) InstMethodOfS() {

}

func (*S) PtrMethodOfS() {
}

type C struct {
	T
	*S
}

func main() {
	var c = C{S: &S{}}
	utils.DumpMethodSet(&c)
	var pc = &C{S: &S{}}
	utils.DumpMethodSet(&pc)
	c.InstMethod1OfT()
	c.PtrMethodOfT()
	c.InstMethodOfS()
	c.PtrMethodOfS()
	pc.InstMethod1OfT()
	pc.PtrMethodOfT()
	pc.InstMethodOfS()
	pc.PtrMethodOfS()
}
