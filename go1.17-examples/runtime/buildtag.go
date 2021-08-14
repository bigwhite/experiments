//go:build linux && !386 && !arm
// +build linux

package main

import "fmt"

func main() {
	fmt.Println("hello, world")
}
