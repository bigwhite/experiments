package main

import (
	"fmt"
	"log"
)

func init() {
	log.Println("plugin1 init")
}

var V int

func F() {
	fmt.Printf("plugin1: public integer variable V=%d\n", V)
}

type foo struct{}

func (foo) M1() {
	fmt.Println("plugin1: invoke foo.M1")
}

var Foo foo
