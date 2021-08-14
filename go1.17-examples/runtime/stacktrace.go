package main

type myStruct struct {
	m int
	s string
	p *float64
}

func foo(a int, b string, c []byte, f *myStruct) (int, error) {
	panic("mypanic")
}

func main() {
	f := 3.14
	ms := myStruct{
		m: 17,
		s: "myStruct",
		p: &f,
	}
	a := 11
	b := "hello"
	c := []byte{'a', 'b', 'c'}
	foo(a, b, c, &ms)
}
