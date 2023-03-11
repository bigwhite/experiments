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

// 测试代码
func TestAdd(t *testing.T) {
	t.Log("g:", curGoroutineID())
	t.Parallel()
	got := Add(2, 3)
	if got != 5 {
		t.Errorf("Add(2, 3) got %d, want 5", got)
	}
}

func TestAddZero(t *testing.T) {
	t.Parallel()
	got := Add(2, 0)
	if got != 2 {
		t.Fatalf("Add(2, 0) got %d, want 2", got)
	}
}

func TestAddOppositeNum(t *testing.T) {
	got := Add(2, -2)
	t.Parallel()
	if got != 0 {
		t.Errorf("Add(2, -2) got %d, want 0", got)
	}
}

func TestAddWithTable(t *testing.T) {
	t.Log("g:", curGoroutineID())
	t.Parallel()
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
		got := Add(caze.a, caze.b)
		if got != caze.r {
			t.Errorf("%s got %d, want %d", caze.name, got, caze.r)
		}
	}
}
