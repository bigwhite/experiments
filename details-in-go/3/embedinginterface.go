package main

import (
	"fmt"

	"./utils"
)

type I1 interface {
	I1Method1()
	I1Method2()
}
type I2 interface {
	I2Method()
}

type I3 interface {
	I1
	I2
}

type T struct {
	I1
}

func (T) Method1() {

}

type I1Impl struct {
}

func (I1Impl) I1Method1() {
	fmt.Println("I1Method1 of I1Impl invoked")
}

func (I1Impl) I1Method2() {
}

func main() {
	utils.DumpMethodSet((*I1)(nil))
	utils.DumpMethodSet((*I2)(nil))
	utils.DumpMethodSet((*I3)(nil))

	var t T
	utils.DumpMethodSet(&t)
	var pt = &T{
		I1: I1Impl{},
	}
	utils.DumpMethodSet(&pt)
	pt.I1Method1()
}
