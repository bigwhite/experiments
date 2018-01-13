package app_c

import (
	"lib_a"
	"lib_b"
)

func main() {
	var b lib_b.B
	lib_a.Foo(b)
}
