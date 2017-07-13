package main

// type definitions
type MyInt int
type MyInt1 MyInt

func main() {
	var i int = 5
	var mi MyInt = 6
	var mi1 MyInt1 = 7

	mi = MyInt(i)
	mi1 = MyInt1(i)
	mi1 = MyInt1(mi)

	mi = i   //Error: cannot use i (type int) as type MyInt in assignment
	mi1 = i  //Error: cannot use i (type int) as type MyInt1 in assignment
	mi1 = mi //Error: cannot use mi (type MyInt) as type MyInt1 in assignment
}
