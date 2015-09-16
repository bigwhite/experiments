package main

import "fmt"

var (
	s = "This is an example about code newline," +
		"for string as right value"
	d = 5 + 4 + 7 +
		4
	a = [...]int{5, 6, 7,
		8}
	m = make(map[string]int,
		100)
	c struct {
		m1     string
		m2, m3 int
		m4     *float64
	}

	f func(int,
		float32) (int,
		error)
)

func foo(int, int) (string, error) {
	return "",
		nil
}

func main() {
	if i := d; i >
		100 {
	}

	var sum int
	for i := 0; i < 100; i = i +
		1 {
		sum += i
	}

	foo(1,
		6)

	var i int
	fmt.Printf("%s, %d\n",
		"this is a demo"+
			" of fmt Printf",
		i)
}
