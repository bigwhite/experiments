/*
 * Copyright (C) iwobi.net. All rights reserved.
 *
 * ringbuf_test.go
 *
 * testcases for ringbuf
 */

package ringbuf

import (
	"iwobi.net/utils/assert"
	"testing"
)

func TestRingBufNew(t *testing.T) {
	rb := New(10)
	assert.AssertNotEquals(t, nil, rb)
	assert.AssertEquals(t, 11, rb.size)
	assert.AssertEquals(t, 0, rb.head)
	assert.AssertEquals(t, 1, rb.tail)
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())
	assert.AssertEquals(t, true, rb.IsEmpty())
}

func TestRingBufSimpleWrite(t *testing.T) {
	rb := New(10)

	n, err := rb.Write([]byte("01234"), 4)

	assert.AssertEquals(t, 4, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 11, rb.size)
	assert.AssertEquals(t, 0, rb.head)
	assert.AssertEquals(t, 5, rb.tail)
	assert.AssertEquals(t, 4, rb.Len())
	assert.AssertEquals(t, 6, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())

	n, err = rb.Write([]byte("456789"), 5)

	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 11, rb.size)
	assert.AssertEquals(t, 9, rb.Len())
	assert.AssertEquals(t, 1, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())

	n, err = rb.Write([]byte("hello"), 5)
	assert.AssertEquals(t, 1, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 11, rb.size)
	assert.AssertEquals(t, 10, rb.Len())
	assert.AssertEquals(t, 0, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, true, rb.IsFull())

	n, err = rb.Write([]byte("world"), 5)
	assert.AssertEquals(t, 0, n)
	assert.AssertEquals(t, errNoSpace, err)
}

func TestRingBufSimpleRead(t *testing.T) {
	rb := New(10)

	var data [64]byte
	n, err := rb.Read(data[:], 5, false)
	assert.AssertEquals(t, 0, n)
	assert.AssertEquals(t, nil, err)

	n, err = rb.Write([]byte("01234"), 5)
	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)

	n, err = rb.Write([]byte("56789"), 5)
	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 10, rb.Len())
	assert.AssertEquals(t, 0, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())

	n, err = rb.Read(data[:], 0, false)
	assert.AssertEquals(t, 0, n)
	assert.AssertEquals(t, nil, err)

	n, err = rb.Read(data[:], 1, false)
	assert.AssertEquals(t, 1, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 1, rb.head)
	assert.AssertEquals(t, 9, rb.Len())
	assert.AssertEquals(t, 1, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "0", string(data[0]))

	n, err = rb.Read(data[:], 4, false)
	assert.AssertEquals(t, 4, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 5, rb.Len())
	assert.AssertEquals(t, 5, rb.Avail())
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "1234", string(data[0:4]))

	n, err = rb.Read(data[:], 5, false)
	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())
	assert.AssertEquals(t, true, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "56789", string(data[0:5]))
}

func TestRingBufComplexWR(t *testing.T) {
	rb := New(10)

	var data [64]byte

	/* h - t - t_n */
	n, err := rb.Write([]byte("01234"), 5)
	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)

	n, err = rb.Read(data[:], 5, false)
	assert.AssertEquals(t, 5, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, true, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "01234", string(data[0:5]))
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())

	/* t_n - h - t */
	n, err = rb.Write([]byte("0123456789"), 10)
	assert.AssertEquals(t, 10, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, true, rb.IsFull())
	assert.AssertEquals(t, 10, rb.Len())
	assert.AssertEquals(t, 0, rb.Avail())

	/* t - t_n - h */
	n, err = rb.Read(data[:], 4, false)
	assert.AssertEquals(t, 4, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "0123", string(data[0:4]))
	assert.AssertEquals(t, 6, rb.Len())
	assert.AssertEquals(t, 4, rb.Avail())

	n, err = rb.Read(data[:], 6, false)
	assert.AssertEquals(t, 6, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, true, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "456789", string(data[0:6]))
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())

	/* h - h_n - t */
	n, err = rb.Read(data[:], 3, false)
	assert.AssertEquals(t, 0, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, true, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())

	/* h_n - t - h */
	n, err = rb.Write([]byte("0123456789"), 10)
	assert.AssertEquals(t, 10, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, true, rb.IsFull())
	assert.AssertEquals(t, 10, rb.Len())
	assert.AssertEquals(t, 0, rb.Avail())

	/* t - h - h_n */
	n, err = rb.Read(data[:], 10, false)
	assert.AssertEquals(t, 10, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, true, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "0123456789", string(data[:10]))
	assert.AssertEquals(t, 0, rb.Len())
	assert.AssertEquals(t, 10, rb.Avail())

	n, err = rb.Write([]byte("012345"), 6)
	assert.AssertEquals(t, 6, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, 6, rb.Len())
	assert.AssertEquals(t, 4, rb.Avail())

	n, err = rb.Read(data[:], 4, true)
	assert.AssertEquals(t, 4, n)
	assert.AssertEquals(t, nil, err)
	assert.AssertEquals(t, false, rb.IsEmpty())
	assert.AssertEquals(t, false, rb.IsFull())
	assert.AssertEquals(t, "0123", string(data[0:4]))
	assert.AssertEquals(t, 6, rb.Len())
	assert.AssertEquals(t, 4, rb.Avail())
}

func BenchmarkRingBufWriteRead(b *testing.B) {
	b.StopTimer()
	rb := New(1024)
	b.StartTimer()
	var data [512]byte

	for i := 0; i < b.N; i++ {
		rb.Write(data[:], 128)
		rb.Read(data[:], 128, false)
	}
}
