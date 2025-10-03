package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	name1 := "José"
	name2 := "Jose\u0301"
	fmt.Println(name1 == name2) // 输出: false

	// 使用 NFC 形式进行规范化后再比较
	fmt.Println(norm.NFC.String(name1) == norm.NFC.String(name2)) // 输出: true
}
