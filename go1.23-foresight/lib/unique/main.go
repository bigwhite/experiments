package main

import (
	"fmt"
	"unique"
)

func main() {
	// 创建唯一句柄
	s1 := unique.Make("hello")
	s2 := unique.Make("world")
	s3 := unique.Make("hello")

	// s1和s3是相等的，因为它们是同一个字符串值
	fmt.Println(s1 == s3) // true
	fmt.Println(s1 == s2) // false

	// 从句柄获取原始值
	fmt.Println(s1.Value()) // hello
	fmt.Println(s2.Value()) // world
}
