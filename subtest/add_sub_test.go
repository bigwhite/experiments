package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

// 被测代码
func Add(a, b int) int {
	return a + b
}

func TestAddWithSubtest(t *testing.T) {
	//t.Log("g:", curGoroutineID())
	cases := []struct {
		name string
		a    int
		b    int
		r    int
	}{
		{"2+3", 2, 3, 5},
		{"2+0", 2, 0, 2},
		{"2+(-2)", 2, -2, 0},
		//... ...
	}

	for _, caze := range cases {
		t.Run(caze.name, func(t *testing.T) {
			t.Log("g:", curGoroutineID())
			got := Add(caze.a, caze.b)
			if got != caze.r {
				t.Errorf("got %d, want %d", got, caze.r)
			}
		})
	}
}
