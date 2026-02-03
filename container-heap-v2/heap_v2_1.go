package main

import (
	"cmp"
	"fmt"

	"github.com/jba/heap" // 提案的参考实现
)

func main() {
	// 创建一个 int 类型的最小堆
	h := heap.New(cmp.Compare[int])

	// 初始化数据
	h.Init([]int{5, 3, 7, 1})

	// 获取并移除最小值
	fmt.Println(h.TakeMin()) // 输出: 1
	fmt.Println(h.TakeMin()) // 输出: 3
}
