package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	fields := []arrow.Field{
		arrow.Field{Name: "name", Type: arrow.BinaryTypes.String},
		arrow.Field{Name: "age", Type: arrow.PrimitiveTypes.Int32},
	}
	structType := arrow.StructOf(fields...)
	sb := array.NewStructBuilder(memory.DefaultAllocator, structType)
	defer sb.Release()

	names := []string{"Alice", "Bob", "Charlie"}
	ages := []int32{25, 30, 35}
	valid := []bool{true, true, true}

	nameBuilder := sb.FieldBuilder(0).(*array.StringBuilder)
	ageBuilder := sb.FieldBuilder(1).(*array.Int32Builder)

	sb.Reserve(len(names))
	nameBuilder.Resize(len(names))
	ageBuilder.Resize(len(names))

	sb.AppendValues(valid)
	nameBuilder.AppendValues(names, valid)
	ageBuilder.AppendValues(ages, valid)

	arr := sb.NewArray().(*array.Struct)
	defer arr.Release()

	bitmaps := arr.NullBitmapBytes()
	fmt.Println(hex.Dump(bitmaps))
	bufs := arr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	nameArr := arr.Field(0).(*array.String)
	bufs = nameArr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	ageArr := arr.Field(1).(*array.Int32)
	bufs = ageArr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	fmt.Println(arr)
}
