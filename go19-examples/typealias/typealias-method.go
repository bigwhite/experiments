package main

type MyInt = int

func (i *MyInt) Increase(a int) {
	*i = *i + MyInt(a)
}

func main() {
	var mi MyInt = 6
	mi.Increase(5)
}
