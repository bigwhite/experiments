// +build trace

package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var mu sync.Mutex
var m = make(map[uint64]int)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func printTrace(id uint64, name, typ string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "\t"
	}
	fmt.Printf("g[%02d]:%s%s%s\n", id, indents, typ, name)
}

func trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	id := getGID()
	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	mu.Lock()
	v := m[id]
	m[id] = v + 1
	mu.Unlock()
	printTrace(id, name, "->", v+1)
	return func() {
		mu.Lock()
		v := m[id]
		m[id] = v - 1
		mu.Unlock()
		printTrace(id, name, "<-", v)
	}
}
