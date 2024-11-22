package foo

import (
	"flag"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	flag.Parse()
	Init()
	m.Run()
}

func TestProcessItems(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := ProcessItems(input)

	// Wait for all goroutines to complete
	time.Sleep(1000 * time.Millisecond)

	// Verify results
	if len(result) != len(input) {
		t.Fatalf("got len=%d, want len=%d", len(result), len(input))
	}

	// Check if results are correct
	for i, v := range input {
		expected := v * 2
		if result[i] != expected {
			t.Errorf("result[%d] = %d, want %d", i, result[i], expected)
		}
	}
}
