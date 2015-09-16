package utils

import (
	"fmt"
	"reflect"
)

// To dump method set of type T, you should pass a pointer to T
// to DumpMethodSet，include interface type.
//
// e.g.
// for interface type I:
//   utils.DumpMethodSet((*I)(nil))
//
// for non-interface type T:
//   var t T
//   utils.DumpMethodSet(&t)
//
// for non-interface type *T:
//   var pt = &T{}
//   utils.DumpMethodSet(&pt)
//
func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method sets:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("\t", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
