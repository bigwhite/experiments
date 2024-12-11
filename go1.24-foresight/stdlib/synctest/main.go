package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个缓存，条目过期时间为 2 秒
	c := NewCache(2*time.Second, func(key string) int {
		fmt.Printf("Creating new item for key: %s\n", key)
		return len(key)
	})

	// 获取缓存条目
	fmt.Println(c.Get("hello")) // 输出: Creating new item for key: hello
	fmt.Println(c.Get("hello")) // 输出: 5

	// 等待 3 秒，使缓存条目过期
	time.Sleep(3 * time.Second)

	// 再次获取缓存条目，此时应创建新条目
	fmt.Println(c.Get("hello")) // 输出: Creating new item for key: hello
	fmt.Println(c.Get("hello")) // 输出: 5
}

