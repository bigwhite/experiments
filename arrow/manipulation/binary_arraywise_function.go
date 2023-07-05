package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/compute"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	data := []int32{5, 10, 0, 25, 2}
	filterMask := []bool{true, false, true, false, true}

	bldr := array.NewInt32Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues(data, nil)
	arr := bldr.NewArray()
	defer arr.Release()

	bldr1 := array.NewBooleanBuilder(memory.DefaultAllocator)
	defer bldr1.Release()
	bldr1.AppendValues(filterMask, nil)
	filterArr := bldr1.NewArray()
	defer filterArr.Release()

	dat, err := compute.Filter(context.Background(), compute.NewDatum(arr),
		compute.NewDatum(filterArr),
		compute.FilterOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}

	arr1, ok := dat.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	fmt.Println(arr1.MakeArray()) // [5 0 2]
}
