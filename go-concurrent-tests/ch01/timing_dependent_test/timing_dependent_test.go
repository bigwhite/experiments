package timing_dependent

import (
	"fmt"
	"testing"
	"time"
)

// 假设我们有一个服务，它会先处理A，再处理B
func serviceThatProcessesInOrder(out chan string) {
	go func() {
		time.Sleep(5 * time.Millisecond) // 模拟处理A
		out <- "A done"
		time.Sleep(10 * time.Millisecond) // 模拟处理B
		out <- "B done"
		close(out)
	}()
}

func TestServiceOrder(t *testing.T) {
	out := make(chan string, 2)
	serviceThatProcessesInOrder(out)

	// 期望在 10ms 内 A 完成
	select {
	case msg := <-out:
		if msg != "A done" {
			t.Errorf("Expected 'A done', got '%s'", msg)
		}
		fmt.Println(msg)
	case <-time.After(10 * time.Millisecond): // 依赖于 A 在 10ms 内完成
		t.Fatal("Timeout waiting for A")
	}

	// 期望在后续 15ms 内 B 完成
	select {
	case msg := <-out:
		if msg != "B done" {
			t.Errorf("Expected 'B done', got '%s'", msg)
		}
		fmt.Println(msg)
	case <-time.After(15 * time.Millisecond): // 依赖于 B 在 A 之后 15ms 内完成
		t.Fatal("Timeout waiting for B")
	}
}
