package main

import (
	"fmt"
	"os"
	"strings"
)

type Result[T any] struct {
	value T
	err   error
	isOk  bool
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, isOk: true}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err, isOk: false}
}

func (r Result[T]) Bind(f func(T) Result[T]) Result[T] {
	if !r.isOk {
		return Err[T](r.err)
	}
	return f(r.value)
}

// 使用示例
func readFile(filename string) Result[string] {
	content, err := os.ReadFile(filename)
	if err != nil {
		return Err[string](err)
	}
	return Ok(string(content))
}

func processContent(content string) Result[string] {
	// 处理内容...
	return Ok(strings.ToUpper(content))
}

func main() {
	result := readFile("input.txt").Bind(processContent)
	fmt.Println(result)
	result = readFile("input1.txt").Bind(processContent)
	fmt.Println(result)
}
