package main

import (
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	ib := array.NewInt64Builder(memory.DefaultAllocator)
	defer ib.Release()

	ib.AppendValues([]int64{1, 2, 3, 4, 5}, nil)
	i1 := ib.NewInt64Array()
	defer i1.Release()

	ib.AppendValues([]int64{6, 7}, nil)
	i2 := ib.NewInt64Array()
	defer i2.Release()

	ib.AppendValues([]int64{8, 9, 10}, nil)
	i3 := ib.NewInt64Array()
	defer i3.Release()

	c := arrow.NewChunked(
		arrow.PrimitiveTypes.Int64,
		[]arrow.Array{i1, i2, i3},
	)
	defer c.Release()

	for _, arr := range c.Chunks() {
		fmt.Println(arr)
	}

	fmt.Println("chunked length =", c.Len())
	fmt.Println("chunked null count=", c.NullN())
}
