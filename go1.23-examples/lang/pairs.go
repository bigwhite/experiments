package main

import "fmt"

// 使用多个类型参数的类型别名
type Pair[T, U any] = struct {
	First  T
	Second U
}

// 使用Pair类型别名
func MakePair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{First: first, Second: second}
}

// 交换Pair中的元素
func SwapPair[T, U any](p Pair[T, U]) Pair[U, T] {
	return Pair[U, T]{First: p.Second, Second: p.First}
}

func main() {
	// 创建一个int和string的Pair
	intStringPair := MakePair(42, "Answer")
	fmt.Printf("Int-String Pair: %+v\n", intStringPair)

	// 创建一个float64和bool的Pair
	floatBoolPair := Pair[float64, bool]{First: 3.14, Second: true}
	fmt.Printf("Float-Bool Pair: %+v\n", floatBoolPair)

	// 使用自定义类型
	type Person struct {
		Name string
		Age  int
	}
	personStringPair := MakePair(Person{Name: "Alice", Age: 30}, "Developer")
	fmt.Printf("Person-String Pair: %+v\n", personStringPair)

	// 交换Pair中的元素
	swappedPair := SwapPair(intStringPair)
	fmt.Printf("Swapped Int-String Pair: %+v\n", swappedPair)

	// 使用类型推断
	inferredPair := MakePair("Hello", 123)
	fmt.Printf("Inferred Pair: %+v\n", inferredPair)
}
