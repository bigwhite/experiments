package main

type I interface {
	f()
	String() string
}

type implOfI struct{}

func (implOfI) f() {}
func (implOfI) String() string {
	return "implOfI"
}

type J interface {
	g()
	String() string
}

type implOfJ struct{}

func (implOfJ) g() {}
func (implOfJ) String() string {
	return "implOfJ"
}

type Foo struct {
	I
	J
}

func (Foo) String() string {
	return "Foo"
}

func main() {
	f := Foo{
		I: implOfI{},
		J: implOfJ{},
	}
	println(f.String())
}
