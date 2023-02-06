package main

func doSth[T comparable](t T) {
}

func main() {
	n := 2
	var i interface{} = n
	doSth(i)
}
