package main

type Foo struct{}
type Bar = Foo

func (f *Foo) Method1() {
}

func (b *Bar) Method2() {
}
func main() {
	var b Bar
	b.Method1()

	var f Foo
	f.Method2()
}
