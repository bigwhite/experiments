package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	bldr := array.NewFloat32Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues([]float32{1.0, 2.0}, nil)
	bldr.AppendNull()
	bldr.AppendValues([]float32{4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.1}, nil)
	arr := bldr.NewArray()
	defer arr.Release()
	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps))
	bufs := arr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}
	fmt.Println(arr)
}
