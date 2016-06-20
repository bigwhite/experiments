package foo

import (
	"fmt"
	"testing"
)

var tests = []struct {
	A, B int
	Sum  int
}{
	{1, 2, 3},
	{1, 1, 2},
	{2, 1, 3},
}

func assertEqual(A, B, expect int, t *testing.T) {
	if got := A + B; got != expect {
		t.Errorf("got %d; want %d", got, expect)
	}
}

/*
func TestSumSubTest1(t *testing.T) {
	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.A, "+", tc.B), func(t *testing.T) {
			assertEqual(tc.A, tc.B, tc.Sum, t)
		})
	}
}
*/

func TestSumSubTest(t *testing.T) {
	//setup code ... ...

	for i, tc := range tests {
		t.Run("A=1", func(t *testing.T) {
			if tc.A != 1 {
				t.Skip(i)
			}
			assertEqual(tc.A, tc.B, tc.Sum, t)
		})

		t.Run("A=2", func(t *testing.T) {
			if tc.A != 2 {
				t.Skip(i)
			}
			assertEqual(tc.A, tc.B, tc.Sum, t)
		})
	}

	//teardown code ... ...
}

func TestSumInOldWay(t *testing.T) {
	for _, tc := range tests {
		if got := tc.A + tc.B; got != tc.Sum {
			t.Errorf("%d + %d = %d; want %d", tc.A, tc.B, got, tc.Sum)
		}
	}
}

func TestSumSubTestInParalell(t *testing.T) {
	t.Run("blockgroup", func(t *testing.T) {
		for _, tc := range tests {
			tc := tc
			t.Run(fmt.Sprint(tc.A, "+", tc.B), func(t *testing.T) {
				t.Parallel()
				assertEqual(tc.A, tc.B, tc.Sum, t)
			})
		}
	})
	//teardown code
}
