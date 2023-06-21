package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	bldr := array.NewInt64Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues([]int64{1, 2}, nil)
	bldr.AppendNull()
	bldr.AppendValues([]int64{4, 5, 6, 7, 8, 9, 10}, nil)
	arr := bldr.NewArray()
	defer arr.Release()
	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps)) // fb 03 00 00
	fmt.Println(arr)               // [1 2 (null) 4 5 6 7 8 9 10]
}
