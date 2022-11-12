package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var rchs []chan int
	for i := 0; i < 10; i++ {
		rchs = append(rchs, make(chan int))
	}

	// 创建SelectCase
	var cases = createRecvCases(rchs)

	// 消费者goroutine
	go func() {
		defer wg.Done()
		for {
			chosen, recv, ok := reflect.Select(cases)
			if ok {
				fmt.Printf("recv from channel [%d], val=%v\n", chosen, recv)
				continue
			}
			// one of the channels is closed, exit the goroutine
			fmt.Printf("channel [%d] closed, select goroutine exit\n", chosen)
			return
		}
	}()

	// 生产者goroutine
	go func() {
		defer wg.Done()
		var n int
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		for i := 0; i < 10; i++ {
			n = r.Intn(10)
			rchs[n] <- n
		}
		close(rchs[n])
	}()

	wg.Wait()
}

func createRecvCases(rchs []chan int) []reflect.SelectCase {
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
