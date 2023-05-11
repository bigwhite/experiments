package main

import (
	"fmt"
	"reflect"
)

func add(a, b int) int {
	return a + b
}

func main() {
	// 获取函数类型变量
	val := reflect.ValueOf(add)
	// 准备函数参数
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}
	// 调用函数
	result := val.Call(args)
	fmt.Println(result[0].Int()) // 输出：3
}
