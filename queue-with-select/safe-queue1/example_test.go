package queue

import (
	"fmt"
)

func ExampleSafeQueueUsage() {
	var q = NewSafe()
	q.Push(1)
	q.Push(2)

	fmt.Printf("%d ", q.Len())
	v, _ := q.Pop()
	fmt.Printf("%d ", v)

	v, _ = q.Pop()
	fmt.Printf("%d ", v)
	fmt.Printf("%d", q.Len())
	// Output:
	// 2 1 2 0
}
