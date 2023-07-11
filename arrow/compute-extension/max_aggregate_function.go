package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/compute"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	data := []int64{5, 10, 0, 25, 2, 35, 7, 15}
	bldr := array.NewInt64Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues(data, nil)
	arr := bldr.NewArray()
	defer arr.Release()

	dat, err := compute.Max(context.Background(), compute.NewDatum(arr))
	if err != nil {
		fmt.Println(err)
		return
	}

	ad, ok := dat.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	arr1 := ad.MakeArray()
	//arr1, err := ad.ToScalar()
	if err != nil {
		panic(err)
	}
	/*
		bufs := arr1.Data().Buffers()
		for _, buf := range bufs {
			if buf != nil {
				fmt.Println(hex.Dump(buf.Bytes()))
			}
		}
	*/
	fmt.Println(arr1)
}
