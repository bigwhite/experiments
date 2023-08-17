//go:debug panicnil=1

package main

import "fmt"

func foo() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recover panic from", e)
			return
		}
		fmt.Println("panic is nil")
	}()

	panic(nil)
}

func main() {
	foo()
}
