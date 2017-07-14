package mylib

import "testing"

func TestAdd(t *testing.T) {
	a, b := 3, 4
	c := Add(a, b)
	if c != 7 {
		t.Fatal("expected: %d, we got: %d", 7, c)
	}
}
