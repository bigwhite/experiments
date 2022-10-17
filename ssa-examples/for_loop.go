package main

func sumN(n int) int {
	var r int
	for i := 1; i <= n; i++ {
		r = r + i
	}
	return r
}

func main() {
	println(sumN(10))
}
