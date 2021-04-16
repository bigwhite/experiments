package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Foo struct {
	Name string
	age  int
}

func main() {
	// 数组
	var a = [5]int{1, 2, 3, 4, 5}
	vaa := reflect.ValueOf(&a) // reflect Value of Address of arr
	va := vaa.Elem()
	va0 := va.Index(0)
	fmt.Printf("a0 = [%d], va0 = [%d]\n", a[0], va0.Int()) // a0 = [1], va0 = [1]
	va0.SetInt(100 + 1)
	fmt.Printf("after set, a0 = [%d]\n", a[0]) // after set, a0 = [101]

	// 切片
	var s = []int{11, 12, 13}
	vs := reflect.ValueOf(s)
	vs0 := vs.Index(0)
	fmt.Printf("s0 = [%d], vs0 = [%d]\n", s[0], vs0.Int()) // s0 = [11], vs0 = [11]
	vs0.SetInt(100 + 11)
	fmt.Printf("after set, s0 = [%d]\n", s[0]) // after set, s0 = [111]

	// map
	var m = map[int]string{
		1: "tom",
		2: "jerry",
		3: "lucy",
	}

	vm := reflect.ValueOf(m)
	vm_1_v := vm.MapIndex(reflect.ValueOf(1))                      // the reflect Value of the value of key 1
	fmt.Printf("m_1 = [%s], vm_1 = [%s]\n", m[1], vm_1_v.String()) // m_1 = [tom], vm_1 = [tom]
	vm.SetMapIndex(reflect.ValueOf(1), reflect.ValueOf("tony"))
	fmt.Printf("after set, m_1 = [%s]\n", m[1]) // after set, m_1 = [tony]

	// 为map m新增一组key-value
	vm.SetMapIndex(reflect.ValueOf(4), reflect.ValueOf("amy"))
	fmt.Printf("after set, m = [%#v]\n", m) // after set, m = [map[int]string{1:"tony", 2:"jerry", 3:"lucy", 4:"amy"}]

	// 结构体
	var f = Foo{
		Name: "lily",
		age:  16,
	}

	vaf := reflect.ValueOf(&f)
	vf := vaf.Elem()
	field1 := vf.FieldByName("Name")
	fmt.Printf("the Name of f = [%s]\n", field1.String()) // the Name of f = [lily]
	field2 := vf.FieldByName("age")
	fmt.Printf("the age of f = [%d]\n", field2.Int()) // the age of f = [16]

	field1.SetString("ally")
	// field2.SetInt(8) // panic: reflect: reflect.Value.SetInt using value obtained using unexported field
	nAge := reflect.NewAt(field2.Type(), unsafe.Pointer(field2.UnsafeAddr())).Elem()
	nAge.SetInt(8)
	fmt.Printf("after set, f is [%#v]\n", f) // after set, f is [main.Foo{Name:"ally", age:8}]

	// 接口
	var g = Foo{
		Name: "Jordan",
		age:  40,
	}

	// 接口底层动态类型为复合类型变量
	var i interface{} = &g
	vi := reflect.ValueOf(i)
	vg := vi.Elem()

	field1 = vg.FieldByName("Name")
	fmt.Printf("the Name of g = [%s]\n", field1.String()) // the Name of g = [Jordan]
	field2 = vg.FieldByName("age")
	fmt.Printf("the age of g = [%d]\n", field2.Int()) // the age of g = [40]

	nAge = reflect.NewAt(field2.Type(), unsafe.Pointer(field2.UnsafeAddr())).Elem()
	nAge.SetInt(50)
	fmt.Printf("after set, g is [%#v]\n", g) // after set, g is [main.Foo{Name:"Jordan", age:50}]

	// 接口底层动态类型为基本类型变量
	var n = 5
	i = &n
	vi = reflect.ValueOf(i).Elem()
	fmt.Printf("i = [%d], vi = [%d]\n", n, vi.Int()) // i = [5], vi = [5]
	vi.SetInt(10)
	fmt.Printf("after set, n is [%d]\n", n) // after set, n is [10]

	// channel
	var ch = make(chan int, 100)
	vch := reflect.ValueOf(ch)
	vch.Send(reflect.ValueOf(22))

	j := <-ch
	fmt.Printf("recv [%d] from channel\n", j) // recv [22] from channel

	ch <- 33

	vj, ok := vch.Recv()
	fmt.Printf("recv [%d] ok[%t]\n", vj.Int(), ok) // recv [33] ok[true]
}
