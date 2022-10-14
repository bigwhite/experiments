package main

func g() {
}
func h() {
}
func f(n int) {
	if n == 0 {
		return
	}
	g()
	h()
	f(n - 1)
}

func main() {
	f(5)
}
