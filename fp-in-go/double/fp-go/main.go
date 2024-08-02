package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/ioeither"
)

// 读取文件内容
func readFile(filename string) ioeither.IOEither[error, string] {
	return ioeither.TryCatchError(func() (string, error) {
		content, err := os.ReadFile(filename)
		return string(content), err
	})
}

// 将字符串转换为数字列表
func parseNumbers(content string) either.Either[error, []int] {
	numbers := []int{}
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return either.Left[[]int](err)
		}
		numbers = append(numbers, num)
	}
	return either.Right[error](numbers)
}

// 将数字乘以2
func multiplyBy2(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		result[i] = num * 2
	}
	return result
}

// 将结果写入文件
func writeFile(filename string, content string) ioeither.IOEither[error, string] {
	return ioeither.TryCatchError(func() (string, error) {
		return "", os.WriteFile(filename, []byte(content), 0644)
	})
}

func main() {
	program := ioeither.Chain(func(content string) ioeither.IOEither[error, string] {
		return ioeither.FromEither(
			either.Chain(func(numbers []int) either.Either[error, string] {
				multiplied := multiplyBy2(numbers)
				result := []string{}
				for _, num := range multiplied {
					result = append(result, strconv.Itoa(num))
				}
				return either.Of[error](strings.Join(result, "\n"))
			})(parseNumbers(content)),
		)
	})(readFile("input.txt"))

	program = ioeither.Chain(func(content string) ioeither.IOEither[error, string] {
		return writeFile("output.txt", content)
	})(program)

	result := program()
	err := either.ToError(result)

	if err != nil {
		fmt.Println("Program failed:", err)
	} else {
		fmt.Println("Program completed successfully")
	}
}
