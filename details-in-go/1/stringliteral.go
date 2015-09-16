package main

import "fmt"

var s = "中国人"
var s1 = "\u4e2d\u56fd\u4eba"
var s2 = "\U00004e2d\U000056fd\U00004eba"
var s3 = "\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba"

func main() {
	fmt.Println(s)
	fmt.Println(`\u4e2d\u56fd\u4eba -> `, s1)
	fmt.Println(`\U00004e2d\U000056fd\U00004eba -> `, s2)
	fmt.Println(`\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba ->`, s3)

	fmt.Println("s byte sequence:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Println("")

	fmt.Println("s1 byte sequence:")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("0x%x ", s1[i])
	}
	fmt.Println("")
}
