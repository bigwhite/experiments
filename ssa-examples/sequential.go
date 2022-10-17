package main

func sum(a, b, c int) int {
	d := a + b
	e := d + c
	return e
}

func main() {
	println(sum(1, 2, 3))
}
