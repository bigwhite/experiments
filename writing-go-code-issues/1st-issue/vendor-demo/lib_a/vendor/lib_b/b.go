package lib_b

import "fmt"

type B struct {
}

func (*B) Do() {
	fmt.Println("lib_b version:v0.1")
}
