package main

func A1() {
	defer trace()()
	B1()
}

func B1() {
	defer trace()()
	C1()
}

func C1() {
	defer trace()()
	D()
}

func D() {
	defer trace()()
}

func main() {
	A1()
}
