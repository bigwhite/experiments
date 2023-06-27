package main

import (
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	schema := arrow.NewSchema(
		[]arrow.Field{
			{Name: "archer", Type: arrow.BinaryTypes.String},
			{Name: "location", Type: arrow.BinaryTypes.String},
			{Name: "year", Type: arrow.PrimitiveTypes.Int16},
		},
		nil,
	)

	rb := array.NewRecordBuilder(memory.DefaultAllocator, schema)
	defer rb.Release()

	rb.Field(0).(*array.StringBuilder).AppendValues([]string{"tony", "amy", "jim"}, nil)
	rb.Field(1).(*array.StringBuilder).AppendValues([]string{"beijing", "shanghai", "chengdu"}, nil)
	rb.Field(2).(*array.Int16Builder).AppendValues([]int16{1992, 1993, 1994}, nil)

	rec := rb.NewRecord()
	defer rec.Release()

	fmt.Println(rec)

	sl := rec.NewSlice(0, 2)
	fmt.Println(sl)
	cols := sl.Columns()
	a1 := cols[0]
	fmt.Println(a1)
}
