// Copyright (C) iWobi.net. All rights reserved.
//
// assert.go
//
// assertion utility for unittesting
//
package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func errorPosition(skip int) (file string, line int, ok bool) {
	_, absfile, line, ok := runtime.Caller(skip)
	if ok {
		file = absfile[strings.LastIndex(absfile, "/")+1:]
	}

	return
}

func equal(expected, actual interface{}) bool {
	v1 := reflect.ValueOf(expected)
	v2 := reflect.ValueOf(actual)

	// both expected and actual are nil
	if v1.IsValid() == false && v2.IsValid() == false {
		return true
	} else {
		// one of them is nil, the other is null reference or null pointer
		// remember: interface{} value is a box, it does not equals nil directly
		// two nil pointers/ref of different type are not equal
		if (v1.IsValid() == false && v2.IsNil()) ||
			(v2.IsValid() == false && v1.IsNil()) {
			return true
		}
	}

	return reflect.DeepEqual(expected, actual)
}

func AssertEquals(t *testing.T, expected, actual interface{}) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(expected, actual); !result {
		fmt.Printf("%s:%d: AssertEqual Error: expect value is <%v> of type <%T>, "+
			"but actual value is <%v> of type <%T>\n",
			file, line, expected, expected, actual, actual)
		t.Fail()
	}
}

func AssertNotEquals(t *testing.T, nexpected, actual interface{}) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(nexpected, actual); result {
		fmt.Printf("%s:%d: AssertNotEqual Error: the actual value is <%v> of type <%T>, it is equal to"+
			" the value <%v> of type <%T> which we do not expect\n",
			file, line, actual, actual, nexpected, nexpected)
		t.Fail()
	}
}

func AssertTrue(t *testing.T, condition bool) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(true, condition); !result {
		fmt.Printf("%s:%d: AssertTrue Error: the actual value is <%v> of type <%T>, "+
			"it is not the value true we expected!\n",
			file, line, condition, condition)
		t.Fail()
	}
}

func AssertFalse(t *testing.T, condition bool) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(false, condition); !result {
		fmt.Printf("%s:%d: AssertTrue Error: the actual value is <%v> of type <%T>, "+
			"it is not the value false we expected!\n", file, line, condition, condition)
		t.Fail()
	}
}

func AssertNil(t *testing.T, actual interface{}) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(nil, actual); !result {
		fmt.Printf("%s:%d: AssertNil Error: the actual value is <%v> of type <%T>, "+
			"it is not the value nil we expected!\n", file, line, actual, actual)
		t.Fail()
	}
}

func AssertNotNil(t *testing.T, actual interface{}) {
	file, line, ok := errorPosition(2)
	if !ok {
		t.Error("Runtime error!")
	}

	if result := equal(nil, actual); result {
		fmt.Printf("%s:%d: AssertNotNil Error: the actual value is <%v> of type <%T>, "+
			"it equals the value nil we do not expect!\n", file, line, actual, actual)
		t.Fail()
	}
}
