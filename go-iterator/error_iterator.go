package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Lines 返回一个迭代器，用于逐行读取 io.Reader 的内容
// 使用 bufio.Reader.ReadLine() 来读取每一行并处理错误
func Lines(r io.Reader) func(func(string, error) bool) {
	br := bufio.NewReader(r)
	return func(yield func(string, error) bool) {
		for {
			line, isPrefix, err := br.ReadLine()
			if err != nil {
				// 如果是 EOF，我们不将其视为错误
				if err != io.EOF {
					yield("", err)
				}
				return
			}

			// 如果一行太长，isPrefix 会为 true，我们需要继续读取
			fullLine := string(line)
			for isPrefix {
				line, isPrefix, err = br.ReadLine()
				if err != nil {
					yield(fullLine, err)
					return
				}
				fullLine += string(line)
			}

			if !yield(fullLine, nil) {
				return
			}
		}
	}
}

func main() {
	reader := strings.NewReader("Hello\nWorld\nGo 1.23\nThis is a very long line that might exceed the buffer size")

	for line, err := range Lines(reader) {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Println(line)
	}
}
