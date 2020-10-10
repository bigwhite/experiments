package main

func foo() {
	var i uint32 = 17
	panic(i)
}

type myint uint32

func bar() {
	var i myint = 27
	panic(i)
}

func main() {
	//foo()
	bar()
}
