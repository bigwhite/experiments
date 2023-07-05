package main

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/compute"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

func main() {
	data1 := []int32{5, 10, 0, 25, 2}
	data2 := []int32{1, 5, 2, 10, 5}
	scalarData1 := int32(6)

	bldr := array.NewInt32Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues(data1, nil)
	arr1 := bldr.NewArray()
	defer arr1.Release()

	bldr.AppendValues(data2, nil)
	arr2 := bldr.NewArray()
	defer arr2.Release()

	result1, err := compute.Add(context.Background(), compute.ArithmeticOptions{},
		compute.NewDatum(arr1),
		compute.NewDatum(arr2))
	if err != nil {
		fmt.Println(err)
		return
	}

	result2, err := compute.Add(context.Background(), compute.ArithmeticOptions{},
		compute.NewDatum(arr1),
		compute.NewDatum(scalarData1))
	if err != nil {
		fmt.Println(err)
		return
	}

	resultArr1, ok := result1.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	fmt.Println(resultArr1.MakeArray()) // [6 15 2 35 7]

	resultArr2, ok := result2.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	fmt.Println(resultArr2.MakeArray()) // [11 16 6 31 8]
}
