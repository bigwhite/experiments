package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 整型
	var i int = 11
	vi := reflect.ValueOf(i)                         // reflect Value of i
	fmt.Printf("i = [%d], vi = [%d]\n", i, vi.Int()) // i = [11], vi = [11]
	// vi.SetInt(11 + 100) // panic: reflect: reflect.Value.SetInt using unaddressable value

	vai := reflect.ValueOf(&i) // reflect Value of Address of i
	vi = vai.Elem()
	fmt.Printf("i = [%d], vi = [%d]\n", i, vi.Int()) // i = [11], vi = [11]
	vi.SetInt(11 + 100)
	fmt.Printf("after set, i = [%d]\n", i) // after set, i = [111]

	// 整型指针
	i = 11
	var pi *int = &i
	vpi := reflect.ValueOf(pi) // reflect Value of pi
	vi = vpi.Elem()
	vi.SetInt(11 + 100)
	fmt.Printf("after set, i = [%d]\n", i) // after set, i = [111]

	// 浮点型
	var f float64 = 3.1415

	vaf := reflect.ValueOf(&f)
	vf := vaf.Elem()
	fmt.Printf("f = [%f], vf = [%f]\n", f, vf.Float()) // f = [3.141500], vf = [3.141500]
	vf.SetFloat(100 + 3.1415)
	fmt.Printf("after set, f = [%f]\n", f) // after set, f = [103.141500]

	// 复数型
	var c = complex(5.1, 6.2)

	vac := reflect.ValueOf(&c)
	vc := vac.Elem()
	fmt.Printf("c = [%g], vc = [%g]\n", f, vc.Complex()) // c = [103.1415], vc = [(5.1+6.2i)]
	vc.SetComplex(complex(105.1, 106.2))
	fmt.Printf("after set, c = [%g]\n", c) // after set, c = [(105.1+106.2i)]

	// 布尔类型
	var b bool = true

	vab := reflect.ValueOf(&b)
	vb := vab.Elem()
	fmt.Printf("b = [%t], vb = [%t]\n", b, vb.Bool()) // b = [true], vb = [true]
	vb.SetBool(false)
	fmt.Printf("after set, b = [%t]\n", b) // after set, b = [false]

	// 字符串类型
	var s string = "hello, reflect"

	vas := reflect.ValueOf(&s)
	vs := vas.Elem()
	fmt.Printf("s = [%s], vs = [%s]\n", s, vs.String()) // s = [hello, reflect], vs = [hello, reflect]
	vs.SetString("bye, reflect")
	fmt.Printf("after set, s = [%s]\n", s) // after set, s = [bye, reflect]
}
