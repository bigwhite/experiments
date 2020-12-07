package main

import (
	"fmt"
	"runtime"
)

func trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Printf("enter: %s\n", name)
	return func() { fmt.Printf("exit: %s\n", name) } // this will be deferred
}
