package main

func foo2() {
	var i int
	for i = 0; i < 10; i++ {
		longa()
	}
}

func foo1() {
	var i int
	for i = 0; i < 100; i++ {
		longa()
	}
}

func main() {
	foo1()
	foo2()
}
