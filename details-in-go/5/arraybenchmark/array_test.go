package array

import "testing"

func BenchmarkArrayRangeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrayRangeLoop()
	}
}
func BenchmarkPointerToArrayRangeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointerToArrayRangeLoop()
	}
}
func BenchmarkSliceRangeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceRangeLoop()
	}
}
