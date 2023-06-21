package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

var (
	F32 arrow.UnionTypeCode = 7
	I32 arrow.UnionTypeCode = 13
)

func main() {
	childFloat32Bldr := array.NewFloat32Builder(memory.DefaultAllocator)
	childInt32Bldr := array.NewInt32Builder(memory.DefaultAllocator)

	defer func() {
		childFloat32Bldr.Release()
		childInt32Bldr.Release()
	}()

	ub := array.NewSparseUnionBuilderWithBuilders(memory.DefaultAllocator,
		arrow.SparseUnionOf([]arrow.Field{
			{Name: "f32", Type: arrow.PrimitiveTypes.Float32, Nullable: true},
			{Name: "i32", Type: arrow.PrimitiveTypes.Int32, Nullable: true},
		}, []arrow.UnionTypeCode{F32, I32}),
		[]array.Builder{childFloat32Bldr, childInt32Bldr})
	defer ub.Release()

	ub.Append(I32)
	childInt32Bldr.Append(5)
	childFloat32Bldr.AppendEmptyValue()

	ub.Append(F32)
	childFloat32Bldr.Append(1.2)
	childInt32Bldr.AppendEmptyValue()

	ub.AppendNull()

	ub.Append(F32)
	childFloat32Bldr.Append(3.4)
	childInt32Bldr.AppendEmptyValue()

	ub.Append(I32)
	childInt32Bldr.Append(6)
	childFloat32Bldr.AppendEmptyValue()

	arr := ub.NewSparseUnionArray()
	defer arr.Release()

	// print type buffer
	buf := arr.TypeCodes().Buf()
	fmt.Println(hex.Dump(buf))

	// print child
	bufs := arr.Field(0).Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	bufs = arr.Field(1).Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	fmt.Println(arr)
}
