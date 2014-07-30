/*
 * Copyright 2012 iWobi.net. All rights reserved.
 *
 * test cases for proto parser
 */

package proto

import (
	"iwobi.net/utils/assert"
	"iwobi.net/utils/ringbuf"
	"testing"
)

func TestParseEmtpyRingBuf(t *testing.T) {
	rb := ringbuf.New(512)
	packages, err := Parse(rb)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 0, len(packages))
}

func TestParseCompletedPackage(t *testing.T) {
	rb := ringbuf.New(512)
	rb.Write([]byte("incr key 1\r\n"), len("incr key 1\r\n"))
	packages, err := Parse(rb)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 1, len(packages))
	assert.AssertEquals(t, INCR, packages[0].Cmd)
	assert.AssertEquals(t, "key", packages[0].Key)
	assert.AssertEquals(t, "1", packages[0].Val)

	s := "incr key1 10\r\nincr key2 11\r\n"
	rb.Write([]byte(s), len(s))
	packages, err = Parse(rb)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 2, len(packages))
	assert.AssertEquals(t, INCR, packages[0].Cmd)
	assert.AssertEquals(t, "key1", packages[0].Key)
	assert.AssertEquals(t, "10", packages[0].Val)
	assert.AssertEquals(t, INCR, packages[1].Cmd)
	assert.AssertEquals(t, "key2", packages[1].Key)
	assert.AssertEquals(t, "11", packages[1].Val)
}

func TestParseInvalidPackage(t *testing.T) {
	rb := ringbuf.New(512)

	s := "incr key1 10\r\nincr1 key2 11\r\n"
	rb.Write([]byte(s), len(s))
	packages, err := Parse(rb)
	assert.AssertEquals(t, ErrInvalidPack, err)
	assert.AssertEquals(t, 1, len(packages))
}

func TestParseInCompletePackage(t *testing.T) {
	rb := ringbuf.New(512)

	s := "incr key1 10\r\nincr key2"
	rb.Write([]byte(s), len(s))
	packages, err := Parse(rb)
	assert.AssertEquals(t, ErrInComplete, err)
	assert.AssertEquals(t, 1, len(packages))
	assert.AssertEquals(t, INCR, packages[0].Cmd)
	assert.AssertEquals(t, "key1", packages[0].Key)
	assert.AssertEquals(t, "10", packages[0].Val)

	s = " 11\r\n"
	rb.Write([]byte(s), len(s))
	packages, err = Parse(rb)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 1, len(packages))
	assert.AssertEquals(t, INCR, packages[0].Cmd)
	assert.AssertEquals(t, "key2", packages[0].Key)
	assert.AssertEquals(t, "11", packages[0].Val)
}

func BenchmarkPackageParse(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := "incr key1 10\r\nincr1 key2 11\r\n"
		rb := ringbuf.New(64)
		rb.Write([]byte(s), len(s))
		b.StartTimer()
		Parse(rb)
	}
}
