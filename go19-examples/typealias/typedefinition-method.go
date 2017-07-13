package main

// type definitions
type MyInt int
type MyInt1 MyInt

func (i *MyInt) Increase(a int) {
	*i = *i + MyInt(a)
}

func main() {
	var mi MyInt = 6
	var mi1 MyInt1 = 7
	mi.Increase(5)
	mi1.Increase(5)
}
