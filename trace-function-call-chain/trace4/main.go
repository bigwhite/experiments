package main

import (
	"sync"
	"time"
)

func A1() {
	defer trace()()
	B1()
}

func B1() {
	defer trace()()
	C1()
}

func C1() {
	defer trace()()
	D()
}

func D() {
	defer trace()()
}

func A2() {
	defer trace()()
	B2()
}
func B2() {
	defer trace()()
	C2()
}
func C2() {
	defer trace()()
	D()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	time.Sleep(time.Millisecond * 50)
	A1()
	wg.Wait()
}
