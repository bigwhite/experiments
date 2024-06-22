package main

import (
	"iter"
	"slices"
)

func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()

		for {
			v1, ok1 := next()
			if !ok1 {
				return // 序列结束
			}

			v2, ok2 := next()
			if !ok2 {
				// 序列中有奇数个元素，最后一个元素没有配对
				return // 序列结束
			}

			if !yield(v1, v2) {
				return // 如果 yield 返回 false，停止迭代
			}
		}
	}
}

func main() {
	sl := []string{"go", "java", "rust", "zig", "python"}
	for k, v := range Pairs(slices.Values(sl)) {
		println(k, v)
	}
}
