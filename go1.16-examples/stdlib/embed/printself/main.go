package main

import (
	_ "embed"
	"fmt"
)

//go:embed main.go
var src string

func main() {
	fmt.Print(src)
}
