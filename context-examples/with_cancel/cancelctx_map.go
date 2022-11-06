package main

import (
	"context"
	"fmt"
	"time"
)

// 直接使用parent cancelCtx
func f1(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine created by f1 exit")
		}
	}()
}

// 基于parent cancelCtx创建新的cancelCtx
func f2(ctx context.Context) {
	ctx1, _ := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("goroutine created by f2 exit")
		}
	}()
}

// 使用基于parent cancelCtx创建的valueCtx
func f3(ctx context.Context) {
	ctx1 := context.WithValue(ctx, "key3", "value3")
	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("goroutine created by f3 exit")
		}
	}()
}

// 基于parent cancelCtx创建的valueCtx之上创建cancelCtx
func f4(ctx context.Context) {
	ctx1 := context.WithValue(ctx, "key4", "value4")
	ctx2, _ := context.WithCancel(ctx1)
	go func() {
		select {
		case <-ctx2.Done():
			fmt.Println("goroutine created by f4 exit")
		}
	}()
}

func main() {
	valueCtx := context.WithValue(context.Background(), "key0", "value0")
	cancelCtx, cf := context.WithCancel(valueCtx)
	f1(cancelCtx)
	f2(cancelCtx)
	f3(cancelCtx)
	f4(cancelCtx)

	time.Sleep(3 * time.Second)
	fmt.Println("cancel all by main")
	cf()
	time.Sleep(10 * time.Second) // wait for log output
}
