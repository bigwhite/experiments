package main

func foo(b bool) int {
	if b {
		return 2
	}
	return 3
}

func main() {
	println(foo(true))
}
