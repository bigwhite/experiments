package pkga

import (
	"demo/pkgb"
	"fmt"
)

func Foo() string {
	pkgb.Zoo()
	return "Hello from Foo!"
}

func Bar() {
	fmt.Println("This is Bar.")
}
