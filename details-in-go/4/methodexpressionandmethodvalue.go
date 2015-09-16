package main

import "fmt"

type T struct {
	a int
}

func (t T) Get() int       { return t.a }
func (t *T) Set(a int) int { t.a = a; return t.a }

func main() {
	var t T

	t.Set(1)
	fmt.Println(t.Get())

	// Method expression.
	//T.Set(2) //invalid method expression T.Set (needs pointer receiver: (*T).Set)
	(*T).Set(&t, 2)
	fmt.Println(T.Get(t))

	f1 := (*T).Set
	f2 := T.Get
	f1(&t, 3)
	fmt.Println(f2(t))

	// Method value.
	f3 := (&t).Set
	fmt.Printf("%T\n", f3)
	f3(4)
	f4 := t.Get
	fmt.Println(f4())

}
