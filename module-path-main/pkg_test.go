package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	n := Add(5, 6)
	if n != 11 {
		t.Errorf("want 11, got %d\n", n)
	}
}
