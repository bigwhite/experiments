package main

func foo(n int) int {
	switch n {
	case 1:
		return 11
	case 2:
		return 12
	}
	return 0
}

func main() {
	println(foo(2))
}
