package utils

import (
	"fmt"
	"reflect"
)

func DumpMethodSet(i interface{}) {
	fmt.Printf("=====%T 's Method Set: =====\n", i)

	v := reflect.TypeOf(i)
	n := v.NumMethod()

	for j := 0; j < n; j++ {
		fmt.Println(v.Method(j).Name)
	}

	fmt.Printf("=====%T 's Method Set end =====\n\n", i)
}
