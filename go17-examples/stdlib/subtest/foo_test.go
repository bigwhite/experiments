package foo

import (
	"fmt"
	"testing"
)

var tests = []struct {
	A, B int
	Sum  int
}{
	{1, 2, 4},
	{1, 1, 3},
	{2, 1, 4},
}

func TestSumSubTest(t *testing.T) {
	for _, tc := range tests {
		//t.Run(fmt.Sprint(tc.A, "+", tc.B), func(t *testing.T) {
		//s := fmt.Sprint(tc.A, "+", tc.B)
		t.Run("No", func(t *testing.T) {
			if got := tc.A + tc.B; got != tc.Sum {
				t.Errorf("got %d; want %d", got, tc.Sum)
			}
		})
	}
}

func TestSumOrigin(t *testing.T) {
	for _, tc := range tests {
		if got := tc.A + tc.B; got != tc.Sum {
			t.Errorf("got %d; want %d", got, tc.Sum)
		}
	}
}

func TestSumInParalell(t *testing.T) {
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprint(tc.A, "+", tc.B), func(t *testing.T) {
			t.Parallel()
			if got := tc.A + tc.B; got != tc.Sum {
				t.Errorf("got %d; want %d", got, tc.Sum)
			}
		})
	}
}
