package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"time"
)

type FileResource struct {
	file *os.File
}

func NewFileResource(filename string) (*FileResource, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// 使用 AddCleanup 注册清理函数
	fd := file.Fd()
	runtime.AddCleanup(file, func(fd uintptr) {
		fmt.Println("Closing file descriptor:", fd)
		syscall.Close(int(fd))
	}, fd)

	return &FileResource{file: file}, nil
}

func main() {
	fileResource, err := NewFileResource("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// 模拟使用 fileResource
	_ = fileResource
	fmt.Println("File opened successfully")

	// 当 fileResource 不再被引用时，AddCleanup 会自动关闭文件
	fileResource = nil
	runtime.GC() // 强制触发 GC，以便清理 fileResource
	time.Sleep(time.Second * 5)
}
