package safe_queue_test

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

var ErrQueueFull = errors.New("queue full")
var ErrQueueEmpty = errors.New("queue empty")
var ErrTimeout = errors.New("operation timed out")

type SafeQueue[T any] struct {
	items chan T
}

func NewSafeQueue[T any](capacity int) *SafeQueue[T] {
	return &SafeQueue[T]{
		items: make(chan T, capacity),
	}
}

func (q *SafeQueue[T]) Enqueue(ctx context.Context, item T, timeout time.Duration) error {
	select {
	case q.items <- item:
		return nil
	case <-time.After(timeout): // 在 bubble 内，这会使用合成时间
		return ErrTimeout
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (q *SafeQueue[T]) Dequeue(ctx context.Context, timeout time.Duration) (T, error) {
	var zero T
	select {
	case item := <-q.items:
		return item, nil
	case <-time.After(timeout): // 在 bubble 内，这会使用合成时间
		return zero, ErrTimeout
	case <-ctx.Done():
		return zero, ctx.Err()
	}
}

func TestSafeQueue_EnqueueDequeue(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		q := NewSafeQueue[int](1) // 容量为 1

		// 测试成功入队
		err := q.Enqueue(t.Context(), 10, 1*time.Millisecond)
		if err != nil {
			t.Fatalf("Enqueue failed: %v", err)
		}

		// 测试队列满时的超时入队
		// 启动一个 goroutine 尝试入队，它应该会因为队列满而等待 time.After
		enqueueErrChan := make(chan error, 1)
		go func() {
			enqueueErrChan <- q.Enqueue(t.Context(), 20, 5*time.Millisecond)
		}()

		synctest.Wait() // 等待入队 goroutine 阻塞在 select 上

		// 合成时间推进，检查入队操作是否确实超时了
		select {
		case err := <-enqueueErrChan:
			if !errors.Is(err, ErrTimeout) {
				t.Fatalf("Expected enqueue to timeout, but got: %v", err)
			}
		}

		// 此时，我们尝试读取队列，不应超时，因为队列中有元素10
		val, err := q.Dequeue(t.Context(), 1*time.Millisecond)
		if err != nil {
			t.Fatalf("Dequeue failed: %v", err)
		}
		if val != 10 {
			t.Fatalf("Dequeued value = %d, want 10", val)
		}

		// 测试队列空时的超时出队
		_, err = q.Dequeue(t.Context(), 2*time.Millisecond)
		if !errors.Is(err, ErrTimeout) {
			t.Fatalf("Expected dequeue to timeout, but got: %v", err)
		}
	})
}
