// 
// Copyright (C) Neusoft. All rights reserved.
// 
// assert_test.go
// 
// testcases for assert package
//

package assert

import (
    "testing"
)

func TestAssertBoolEquals(t *testing.T) {
    AssertEquals(t, true, true)
    AssertEquals(t, false, false)
}

func TestAssertIntEquals(t *testing.T) {
    AssertEquals(t, 5, 5)
    AssertEquals(t, 0, 0)
    AssertEquals(t, -1, 0-1)
}

func TestAssertUintptrEquals(t *testing.T) {
    var a, b uintptr
    a = 0x12345678
    b = 0x12345678

    AssertEquals(t, a, b)
}

func TestAssertFloat64Equals(t *testing.T) {
    AssertEquals(t, 3.1415926, 3.1415926)
}

func TestAssertComplex64Equals(t *testing.T) {
    x := complex64(-18.3 + 8.9i)
    y := complex64(-18.3 + 8.9i)
    AssertEquals(t, x, y)
}

func TestAssertStringEquals(t *testing.T){
    AssertEquals(t, "hello", "hello")
}

func TestAssertIntArrayEquals(t *testing.T) {
    var a [5]int = [5]int{1, 2, 3, 4, 5}
    var b [5]int = [5]int{1, 2, 3, 4, 5}
    AssertEquals(t, a, b)
}

type foo struct {
    name string
    age int
}

func TestAssertStructEquals(t *testing.T) {
    a := foo{"tony", 20}
    b := foo{"tony", 20}

    AssertEquals(t, a, b)
}

func TestAssertNilEquals(t *testing.T) {
    var p *int
    AssertEquals(t, nil, nil)
    AssertEquals(t, nil, p)
    var q *int
    AssertEquals(t, p, q)
}

func TestAssertRefEquals(t *testing.T) {
    var a chan int
    var b chan int
    AssertEquals(t, a, b)
}

func TestAssertBoolNotEquals(t *testing.T) {
    AssertNotEquals(t, true, false)
    AssertNotEquals(t, false, true)
}

func TestAssertIntNotEquals(t *testing.T) {
    AssertNotEquals(t, 0, 6-5)
}

func TestAssertUintptrNotEquals(t *testing.T) {
    var a, b uintptr
    a = 0x12345678
    b = 0x12345679

    AssertNotEquals(t, a, b)
}

func TestAssertFloat64NotEquals(t *testing.T) {
    AssertNotEquals(t, 3.1415926, 3.1415927)
}

func TestAssertComplex64NotEquals(t *testing.T) {
    x := complex64(-18.3 + 8.9i)
    y := complex64(-17.3 + 8.9i)
    AssertNotEquals(t, x, y)
}

func TestAssertStringNotEquals(t *testing.T){
    AssertNotEquals(t, "hello", "hello1")
}

func TestAssertIntArrayNotEquals(t *testing.T) {
    var a [5]int = [5]int{4, 2, 3, 4, 5}
    var b [5]int = [5]int{1, 2, 3, 4, 5}
    AssertNotEquals(t, a, b)
}

func TestAssertStructNotEquals(t *testing.T) {
    a := foo{"tony", 20}
    b := foo{"tony", 21}

    AssertNotEquals(t, a, b)
}

func TestAssertNilNotEquals(t *testing.T) {
    var a int = 5
    var p *int = &a
    AssertNotEquals(t, nil, p)
    var m *float64
    AssertNotEquals(t, p, m)
}

func TestAssertTrue(t *testing.T) {
    AssertTrue(t, 1==1)
}

func TestAssertFalse(t *testing.T) {
    AssertTrue(t, 1 !=0)
}

func TestAssertNil(t *testing.T) {
    AssertNil(t, nil)
    var p *int
    AssertNil(t, p)
}

func TestAssertNotNil(t *testing.T) {
    a := 1
    var p *int = &a
    AssertNotNil(t, p)
}

func TestAssertRefNotEquals(t *testing.T) {
    var a chan int
    var b chan string
    var c map[string]int
    AssertNotEquals(t, a, b)
    AssertNotEquals(t, a, c)
}
