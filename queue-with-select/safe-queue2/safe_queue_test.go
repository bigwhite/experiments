package queue

import (
	"testing"
)

func BenchmarkParallelQueuePush(b *testing.B) {
	b.ReportAllocs()
	q := NewSafe()
	s := "hello, queue"
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Push(s)
		}
	})
}

func BenchmarkParallelQueuePop(b *testing.B) {
	b.ReportAllocs()
	q := NewSafe()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			q.Pop()
		}
	})
}

func BenchmarkParallelPushBufferredChan(b *testing.B) {
	var c = make(chan interface{}, 1000)
	s := "hello, queue"
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		go func() {
			for {
				_ = <-c
			}
		}()
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c <- s
		}
	})
}

func BenchmarkParallelPopBufferedChan(b *testing.B) {
	var c = make(chan interface{}, 1000)
	s := "hello, queue"
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		go func() {
			for {
				c <- s
			}
		}()
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = <-c
		}
	})
}

func BenchmarkParallelPushUnBufferredChan(b *testing.B) {
	var c = make(chan interface{})
	s := "hello, queue"
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		go func() {
			for {
				_ = <-c
			}
		}()
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c <- s
		}
	})
}

func BenchmarkParallelPopUnBufferedChan(b *testing.B) {
	var c = make(chan interface{})
	s := "hello, queue"
	b.ReportAllocs()
	for n := 0; n < 10; n++ {
		go func() {
			for {
				c <- s
			}
		}()
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = <-c
		}
	})
}
