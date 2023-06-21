package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	var (
		vs = [][]int32{{0, 1}, {2, 3, 4, 5}, {6}, {7, 8, 9}}
	)

	lb := array.NewListBuilder(memory.DefaultAllocator, arrow.PrimitiveTypes.Int32)
	defer lb.Release()

	vb := lb.ValueBuilder().(*array.Int32Builder)
	vb.Reserve(len(vs))

	for _, v := range vs {
		lb.Append(true)
		vb.AppendValues(v[:], nil)
	}

	arr := lb.NewArray().(*array.List)
	defer arr.Release()
	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps))
	bufs := arr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	varr := arr.ListValues().(*array.Int32)
	bufs = varr.Data().Buffers()

	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}
	fmt.Println(arr)
}
