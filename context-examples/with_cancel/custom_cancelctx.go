package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func f1(ctx context.Context) {
	ctx1, _ := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx1.Done():
			fmt.Println("goroutine created by f1 exit")
		}
	}()
}

type myCancelCtx struct {
	context.Context
	done chan struct{}
	err  error
}

func (ctx *myCancelCtx) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *myCancelCtx) Err() error {
	return ctx.err
}

func WithMyCancelCtx(parent context.Context) (context.Context, context.CancelFunc) {
	var myCtx = &myCancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}

	return myCtx, func() {
		myCtx.done <- struct{}{}
		myCtx.err = context.Canceled
	}
}

func main() {
	valueCtx := context.WithValue(context.Background(), "key0", "value0")
	fmt.Println("before f1:", runtime.NumGoroutine())

	myCtx, mycf := WithMyCancelCtx(valueCtx)
	f1(myCtx)
	fmt.Println("after f1:", runtime.NumGoroutine())

	time.Sleep(3 * time.Second)
	mycf()
	time.Sleep(10 * time.Second) // wait for log output
}
