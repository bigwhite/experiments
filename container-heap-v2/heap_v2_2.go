package main

import (
	"cmp"
	"fmt"

	"github.com/jba/heap"
)

type Task struct {
	Priority int
	Name     string
	Index    int // 用于记录在堆中的位置
}

func main() {
	// 1. 创建带索引维护功能的堆
	// 提供一个回调函数：当元素移动时，自动更新其 Index 字段
	h := heap.NewIndexed(
		func(a, b *Task) int { return cmp.Compare(a.Priority, b.Priority) },
		func(t *Task, i int) { t.Index = i },
	)

	task := &Task{Priority: 10, Name: "Fix Bug"}

	// 2. 插入任务
	h.Insert(task)
	fmt.Printf("Inserted task index: %d\n", task.Index) // Index 自动更新为 0

	// 3. 修改优先级
	task.Priority = 1     // 变得更紧急
	h.Changed(task.Index) // 极其高效的 O(log n) 更新

	// 4. 取出最紧急的任务
	top := h.TakeMin()
	fmt.Printf("Top task: %s (Priority %d)\n", top.Name, top.Priority)
}
