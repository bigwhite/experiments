package main

import (
	"context"
	"fmt"
)

func main() {
	myError := fmt.Errorf("%s", "myError")
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(myError)
	fmt.Println(ctx.Err())          // returns context.Canceled
	fmt.Println(context.Cause(ctx)) // returns myError
}
