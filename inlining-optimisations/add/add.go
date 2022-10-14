package main

////go:noinline
func add(a, b int) int {
	return a + b
}

func main() {
	var a, b = 5, 6
	c := add(a, b)
	println(c)
}
