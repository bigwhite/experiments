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
	bldr := array.NewInt32Builder(memory.DefaultAllocator)
	defer bldr.Release()
	bldr.AppendValues(data, nil)
	arr := bldr.NewArray()
	defer arr.Release()

	dat, err := compute.Negate(context.Background(), compute.ArithmeticOptions{}, compute.NewDatum(arr))
	if err != nil {
		fmt.Println(err)
		return
	}

	arr1, ok := dat.(*compute.ArrayDatum)
	if !ok {
		fmt.Println("type assert fail")
		return
	}
	fmt.Println(arr1.MakeArray()) // [-5 -10 0 -25 -2]
}
