package main

import (
	"reflect"
	"testing"
)

func createCases(rchs []chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range rchs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	return cases
}

func BenchmarkSelect(b *testing.B) {
	var c1 = make(chan int)
	var c2 = make(chan int)
	var c3 = make(chan int)

	go func() {
		for {
			c1 <- 1
		}
	}()
	go func() {
		for {
			c2 <- 2
		}
	}()
	go func() {
		for {
			c3 <- 3
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		select {
		case <-c1:
		case <-c2:
		case <-c3:
		}
	}
}

func BenchmarkReflectSelect(b *testing.B) {
	var c1 = make(chan int)
	var c2 = make(chan int)
	var c3 = make(chan int)

	go func() {
		for {
			c1 <- 1
		}
	}()
	go func() {
		for {
			c2 <- 2
		}
	}()
	go func() {
		for {
			c3 <- 3
		}
	}()

	chs := createCases([]chan int{c1, c2, c3})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = reflect.Select(chs)
	}
}
