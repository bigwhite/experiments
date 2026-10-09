package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s! (%s/%s)\n", name, runtime.GOOS, runtime.GOARCH)
}
