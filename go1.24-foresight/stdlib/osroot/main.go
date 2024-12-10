package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Root 访问相对路径
	root, err := os.OpenRoot(".") // 打开当前目录作为根目录
	if err != nil {
		fmt.Println("Error opening root:", err)
		return
	}
	defer root.Close()

	// 尝试访问相对路径 "../passwd"
	file, err := root.Open("../passwd")
	if err != nil {
		fmt.Println("Error opening file with os.Root:", err)
	} else {
		fmt.Println("Successfully opened file with os.Root")
		file.Close()
	}

	// 传统的 os.OpenFile 方式
	// 尝试访问相对路径 "../passwd"
	file2, err := os.OpenFile("../passwd", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file with os.OpenFile:", err)
	} else {
		fmt.Println("Successfully opened file with os.OpenFile")
		file2.Close()
	}
}

