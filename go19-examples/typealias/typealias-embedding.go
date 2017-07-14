package main

type Foo struct{}
type Bar = Foo

type SuperFoo struct {
	Bar
}

func (f *Foo) Method1() {
}

func main() {
	var s SuperFoo
	s.Method1()
}
