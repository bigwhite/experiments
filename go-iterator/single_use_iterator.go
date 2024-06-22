package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Lines 返回一个迭代器，用于逐行读取 io.Reader 的内容
func Lines(r io.Reader) func(func(string) bool) {
	scanner := bufio.NewScanner(r)
	return func(yield func(string) bool) {
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

func main() {
	f, err := os.Open("ref.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	itor := Lines(f)
	println("first loop:\n")

	for v := range itor {
		fmt.Println(v)
	}

	println("\nsecond loop:\n")

	for v := range itor {
		fmt.Println(v)
	}
}
