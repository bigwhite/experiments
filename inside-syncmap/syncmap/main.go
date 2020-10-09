package main

import (
	"fmt"

	sync "github.com/bigwhite/go/sync"
)

type val struct {
	s string
}

func main() {
	var m sync.Map
	fmt.Println("sync.Map init status:")
	m.Dump()
	val1 := &val{"val1"}
	m.Store("key1", val1)
	fmt.Println("\nafter store key1:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2:")
	m.Dump()

	val2 := &val{"val2"}
	m.Store("key2", val2)
	fmt.Println("\nafter store key2:")
	m.Dump()

	val3 := &val{"val3"}
	m.Store("key3", val3)
	fmt.Println("\nafter store key3:")
	m.Dump()

	m.Load("key1")
	fmt.Println("\nafter load key1:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2 2nd:")
	m.Dump()

	m.Load("key2")
	fmt.Println("\nafter load key2 3rd:")
	m.Dump()

	/*
		// 验证update key
		val2_1 := &val{"val2_1"}
		m.Store("key2", val2_1)
		fmt.Println("\nafter update key2(in read, not in dirty):")
		m.Dump()

		val4 := &val{"val4"}
		m.Store("key4", val4)
		fmt.Println("\nafter store key4:")
		m.Dump()

		val4_1 := &val{"val4_1"}
		m.Store("key4", val4_1)
		fmt.Println("\nafter update key4(not in read, in dirty):")
		m.Dump()
	*/

	m.Delete("key2")
	fmt.Println("\nafter delete key2:")
	m.Dump()

	val4 := &val{"val4"}
	m.Store("key4", val4)
	fmt.Println("\nafter store key4:")
	m.Dump()

	m.Delete("key4")
	fmt.Println("\nafter delete key4:")
	m.Dump()

	m.Delete("key1")
	fmt.Println("\nafter delete key1:")
	m.Dump()

	/*
		val5 := &val{"val5"}
		m.Store("key5", val5)
		fmt.Println("\nafter store key5:")
		m.Dump()
	*/

	m.Load("key5")
	fmt.Println("\nafter load key5:")
	m.Dump()

	m.Load("key5")
	fmt.Println("\nafter load key5 2nd:")
	m.Dump()
}
