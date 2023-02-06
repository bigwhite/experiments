package main

func F[T1 any]() {
	type x struct{}
	type y = x
}

func main() {
	F[int]()
}
