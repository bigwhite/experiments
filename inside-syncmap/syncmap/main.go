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
	_, ok := m.Load("key1")
	if !ok {
		fmt.Println("not found the key1")
	} else {
		fmt.Println("found the key1")
	}

	val1 := &val{"val1"}
	m.Store("key1", val1)

	fmt.Println("\nafter store key1:")
	m.Dump()

	m.Load("key1")
	fmt.Println("\nafter load key1:")
	m.Dump()

	val2 := &val{"val2"}
	m.Store("key2", val2)
	val3 := &val{"val3"}
	m.Store("key3", val3)

	fmt.Println("\nafter store key2 and key3:")
	m.Dump()

	m.Load("key1")
	fmt.Println("\nafter load key1 again:")
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

	val4 := &val{"val4"}
	m.Store("key4", val4)
	fmt.Println("\nafter store key4:")
	m.Dump()

	for i := 0; i < 4; i++ {
		m.Load("key4")
		fmt.Printf("\nafter load key4 %d:", i+1)
		m.Dump()
	}

	m.Delete("key1")
	fmt.Println("\nafter delete key1:")
	m.Dump()

	val5 := &val{"val5"}
	m.Store("key5", val5)
	fmt.Println("\nafter store key5:")
	m.Dump()

	m.Load("key1")
	fmt.Println("\nafter load key1 after it has been deleted:")
	m.Dump()

	for i := 0; i < 4; i++ {
		m.Load("key5")
		fmt.Printf("\nafter load key5 %d:", i+1)
		m.Dump()
	}

	val5a := &val{"val5a"}
	m.Store("key5", val5a)
	fmt.Println("\nafter store key5 2nd:")
	m.Dump()

	val6 := &val{"val6"}
	m.Store("key6", val6)
	fmt.Println("\nafter store key6:")
	m.Dump()

	val5b := &val{"val5b"}
	m.Store("key5", val5b)
	fmt.Println("\nafter store key5 3rd:")
	m.Dump()

	m.Delete("key5")
	fmt.Println("\nafter delete key5:")
	m.Dump()
}
