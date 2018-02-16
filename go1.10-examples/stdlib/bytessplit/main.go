package main

import (
	"bytes"
	"fmt"
)

func desc(b []byte) string {
	return fmt.Sprintf("len: %2d | cap: %2d | %q\n", len(b), cap(b), b)
}

func main() {
	text := []byte("Hello, Go1.10 is coming!")
	fmt.Printf("text:  %s", desc(text))

	subslices := bytes.Split(text, []byte(" "))
	fmt.Printf("subslice 0:  %s", desc(subslices[0]))
	fmt.Printf("subslice 1:  %s", desc(subslices[1]))
	fmt.Printf("subslice 2:  %s", desc(subslices[2]))
	fmt.Printf("subslice 3:  %s", desc(subslices[3]))
}
