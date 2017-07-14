package main

import (
	"fmt"

	"github.com/bigwhite/experiments/go19-examples/typealias/mylib"
)

func main() {
	f := &mylib.Foo{5, "Hello"}
	f.String()            // ok
	fmt.Println(f.A, f.B) // ok

	// Error:  f.anotherMethod undefined (cannot refer to unexported field
	// or method mylib.(*foo).anotherMethod)
	f.anotherMethod()
}
