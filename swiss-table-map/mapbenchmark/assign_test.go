package gomapbench

import (
	"strconv"
	"testing"
	"unsafe"
)

func BenchmarkMapAssignGrow(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignGrowInt64, cases...))
}

func BenchmarkMapAssignPreAllocate(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignPreAllocateInt64, cases...))
}

func BenchmarkMapAssignReuse(b *testing.B) {
	b.Run("Int64", runWith(benchmarkMapAssignReuseInt64, cases...))
}

func benchmarkMapAssignGrowInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int32]int)
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
	}
}

func benchmarkMapAssignGrowPointer(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[unsafe.Pointer]int)
		for j := 0; j < n; j++ {
			j := j
			m[unsafe.Pointer(&j)] = j
		}
	}
}

func benchmarkMapAssignGrowInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int64]int)
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
	}
}

func benchmarkMapAssignGrowStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a := make(map[string]int)
		for j := 0; j < n; j++ {
			a[k[j]] = i
		}
	}
}

func benchmarkMapAssignPreAllocateInt32(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int32]int, n)
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
	}
}

func benchmarkMapAssignPreAllocatePointer(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[unsafe.Pointer]int, n)
		for j := 0; j < n; j++ {
			j := j
			m[unsafe.Pointer(&j)] = j
		}
	}
}

func benchmarkMapAssignPreAllocateInt64(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		m := make(map[int64]int, n)
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
	}
}

func benchmarkMapAssignPreAllocateStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]int, n)
		for j := 0; j < n; j++ {
			m[k[j]] = i
		}
	}
}

func benchmarkMapAssignReuseInt32(b *testing.B, n int) {
	m := make(map[int32]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m[int32(j)] = j
		}
		for k := range m {
			delete(m, k)
		}
	}
}

func benchmarkMapAssignReusePointer(b *testing.B, n int) {
	m := make(map[unsafe.Pointer]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			j := j
			m[unsafe.Pointer(&j)] = j
		}
		for k := range m {
			delete(m, k)
		}
	}
}

func benchmarkMapAssignReuseInt64(b *testing.B, n int) {
	m := make(map[int64]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m[int64(j)] = j
		}
		for k := range m {
			delete(m, k)
		}
	}
}

func benchmarkMapAssignReuseStr(b *testing.B, n int) {
	k := make([]string, n)
	for i := 0; i < len(k); i++ {
		k[i] = strconv.Itoa(i)
	}
	m := make(map[string]int, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			m[k[j]] = i
		}
		for k := range m {
			delete(m, k)
		}
	}
}
