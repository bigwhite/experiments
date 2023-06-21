package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	dictType := &arrow.DictionaryType{IndexType: &arrow.Int8Type{}, ValueType: &arrow.StringType{}}
	bldr := array.NewDictionaryBuilder(memory.DefaultAllocator, dictType)
	defer bldr.Release()

	bldr.AppendValueFromString("foo")
	bldr.AppendValueFromString("bar")
	bldr.AppendValueFromString("foo")
	bldr.AppendValueFromString("bar")
	bldr.AppendNull()
	bldr.AppendValueFromString("baz")

	arr := bldr.NewDictionaryArray()
	defer arr.Release()
	bufs := arr.Data().Buffers()
	for _, buf := range bufs {
		fmt.Println(hex.Dump(buf.Buf()))
	}

	dict := arr.Dictionary()
	// print value string in dict
	bufs = dict.Data().Buffers()
	for _, buf := range bufs {
		if buf == nil {
			continue
		}
		fmt.Println(hex.Dump(buf.Buf()))
	}

	fmt.Println(arr)
}
