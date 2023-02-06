package main

func doSth[T comparable](t1, t2 T) {
	if t1 != t2 {
		println("unequal")
		return
	}
	println("equal")
}

func main() {
	n1 := []byte{2}
	n2 := []byte{3}
	var i interface{} = n1
	var j interface{} = n2
	doSth(i, j)
}
