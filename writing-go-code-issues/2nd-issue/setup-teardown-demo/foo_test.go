package foo_test

import (
	"fmt"
	"testing"
)

func setUp(t *testing.T, args ...interface{}) func() {
	fmt.Println("testcase setUp")
	// use t and args

	return func() {
		// use t
		// use args
		fmt.Println("testcase tearDown")
	}
}

func TestXXX(t *testing.T) {
	defer setUp(t)()
	fmt.Println("invoke testXXX")
}
