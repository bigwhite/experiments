package main

import (
	"testing"
	"testing/quick"
)

func Foo(a, b int) int {
	return a + b*5
}

func Bar(a, b int) int {
	return a + b*4
}

func TestFoo(t *testing.T) {
	err := quick.CheckEqual(Foo, Bar, nil)
	if err != nil {
		t.Error(err)
	}
}
