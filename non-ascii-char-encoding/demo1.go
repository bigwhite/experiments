package main

import (
	"fmt"

	utils "github.com/bigwhite/gocmpp/utils"
)

func main() {
	var stringLiteral = "中国人"
	var stringUsingRuneLiteral = "\u4E2D\u56FD\u4EBA"

	if stringLiteral != stringUsingRuneLiteral {
		fmt.Println("stringLiteral is not equal to stringUsingRuneLiteral")
		return
	}
	fmt.Println("stringLiteral is equal to stringUsingRuneLiteral")

	for i, v := range stringLiteral {
		fmt.Printf("中文字符: %s <=> Unicode码点(rune): %X <=> UTF8编码(内存值): ", string(v), v)
		s := stringLiteral[i : i+3]
		for _, v := range []byte(s) {
			fmt.Printf("0x%X ", v)
		}

		s1, _ := utils.Utf8ToGB18030(s)
		fmt.Printf("<=> GB18030编码(内存值): ")
		for _, v := range []byte(s1) {
			fmt.Printf("0x%X ", v)
		}
		fmt.Printf("\n")
	}
}
