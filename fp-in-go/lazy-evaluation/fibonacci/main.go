package main

import (
	"fmt"
)

// Fibonacci 返回一个生成无限斐波那契数列的函数
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	fib := Fibonacci()
	for i := 0; i < 10; i++ { // 打印前 10 个斐波那契数
		fmt.Println(fib())
	}
}
