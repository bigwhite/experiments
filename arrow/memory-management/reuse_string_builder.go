package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	bldr := array.NewStringBuilder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues([]string{"hello", "apache arrow"}, nil)
	arr := bldr.NewArray()
	defer arr.Release()
	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps))
	bufs := arr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}
	fmt.Println(arr)

	// reuse the builder
	bldr.AppendValues([]string{"happy birthday", "leo messi"}, nil)
	arr1 := bldr.NewArray()
	defer arr1.Release()
	bitmaps1 := arr1.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps1))
	bufs1 := arr1.Data().Buffers()
	for _, buf := range bufs1 {
		if buf != nil {
			fmt.Println(hex.Dump(buf.Buf()))
		}
	}
	fmt.Println(arr1)
}
