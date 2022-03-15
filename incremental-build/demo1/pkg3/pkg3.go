package pkg3

import "demo1/pkg4"

func F3() {
	println("pkg3: F3")
	pkg4.F4()
}
