package main

func main() {
	f(100)
}

func f(x int) {
	if x < 0 {
		return
	}
	g(x - 1)
}
func g(x int) {
	h(x - 1)
}
func h(x int) {
	f(x - 1)
}
