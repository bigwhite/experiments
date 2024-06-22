package main

import (
	"slices"
	"testing"
)

var sl = []string{"go", "java", "rust", "zig", "python"}

func iterateUsingClassicLoop() {
	for i, v := range sl {
		_, _ = i, v
	}
}

func iterateUsingIterator() {
	for i, v := range slices.All(sl) {
		_, _ = i, v
	}
}

func BenchmarkIterateUsingClassicLoop(b *testing.B) {
	for range b.N {
		iterateUsingClassicLoop()
	}
}

func BenchmarkIterateUsingIterator(b *testing.B) {
	for range b.N {
		iterateUsingIterator()
	}
}
