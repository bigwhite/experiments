package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch0, ch1, ch2 := make(chan int), make(chan int), make(chan int)
	var schs = []chan int{ch0, ch1, ch2}

	// 创建SelectCase
	var cases = createCases(schs)

	// 生产者goroutine
	go func() {
		defer wg.Done()
		for range cases {
			chosen, _, _ := reflect.Select(cases)
			fmt.Printf("send to channel [%d], val=%v\n", chosen, cases[chosen].Send)
			cases[chosen].Chan = reflect.Value{}
		}
		fmt.Println("select goroutine exit")
		return
	}()

	// 消费者goroutine
	go func() {
		defer wg.Done()
		for range schs {
			var v int
			select {
			case v = <-ch0:
				fmt.Printf("recv %d from ch0\n", v)
			case v = <-ch1:
				fmt.Printf("recv %d from ch1\n", v)
			case v = <-ch2:
				fmt.Printf("recv %d from ch2\n", v)
			}
		}
	}()

	wg.Wait()
}

func createCases(schs []chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建send case
	for i, ch := range schs {
		n := i + 100
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(n),
		})
	}

	return cases
}
