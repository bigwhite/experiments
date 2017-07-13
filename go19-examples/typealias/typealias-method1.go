package main

type Foo struct{}
type Bar = Foo

func (f *Foo) Method1() {
}

func main() {
	var b Bar
	b.Method1()
}
