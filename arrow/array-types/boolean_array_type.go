package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	bldr := array.NewBooleanBuilder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues([]bool{true, false}, nil)
	bldr.AppendNull()
	bldr.AppendValues([]bool{true, true, true, false, false, false, true}, nil)
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
