package main

import (
	"demo/pkg1"
	"demo/pkg2"
	"fmt"
)

func main() {
	var a any = 5
	fmt.Println(a)

	pkg1.F1()
	pkg2.F2()
}
