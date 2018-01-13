package lib_a

import "lib_b"

func Foo(b lib_b.B) {
	b.Do()
}
