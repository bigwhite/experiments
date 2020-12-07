// +build trace

package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	id := getGID()
	fmt.Printf("g[%02d]: enter %s\n", id, name)
	return func() { fmt.Printf("g[%02d]: exit %s\n", id, name) } // this will be deferred
}
