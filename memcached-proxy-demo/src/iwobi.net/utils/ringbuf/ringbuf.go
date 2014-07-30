/*
 * Copyright (C) iWobi.net. All rights reserved.
 *
 * ringbuf.go
 *
 * a ringbuf implementation in go
 *
 * ---------------------------
 * head| item1| item2|..|tail|
 * ---------------------------
 */

package ringbuf

import (
	"errors"
)

var errNoSpace = errors.New("ringbuf has no space")

type RingBuf struct {
	data []byte
	size int
	head int
	tail int
}

func New(cap int) (r *RingBuf) {
	r = new(RingBuf)

	r.size = cap + 1
	r.data = make([]byte, r.size)

	/*
	 * head is a pilot index
	 * buff full: head == tail
	 * buff empty: (head + cap + 1) % cap == tail
	 */
	r.head = 0
	r.tail = r.head + 1

	return
}

func (r *RingBuf) Len() int {
	return (r.tail - r.head - 1 + r.size) % r.size
}

func (r *RingBuf) Avail() int {
	return r.size - 1 - r.Len()
}

func (r *RingBuf) IsEmpty() bool {
	return r.Len() == 0
}

func (r *RingBuf) IsFull() bool {
	return r.Avail() == 0
}

func (r *RingBuf) Write(data []byte, n int) (int, error) {
	if n == 0 {
		return 0, nil
	}

	if r.Avail() == 0 {
		return 0, errNoSpace
	}

	writeLen := n
	if n > r.Avail() {
		writeLen = r.Avail()
	}
	nextTail := (r.tail + writeLen) % r.size

	if nextTail > r.tail {
		copy(r.data[r.tail:], data)
	} else {
		left := r.size - r.tail
		copy(r.data[r.tail:], data[:left])
		copy(r.data[0:], data[left:])
	}

	r.tail = nextTail
	return writeLen, nil
}

func (r *RingBuf) Read(data []byte, n int, peek bool) (int, error) {
	if n == 0 || r.Len() == 0 {
		return 0, nil
	}

	toReadLen := r.Len()
	if toReadLen > n {
		toReadLen = n
	}

	headNext := (r.head + toReadLen) % r.size

	if headNext > r.head {
		copy(data[0:], r.data[(r.head+1):(r.head+1+toReadLen)])
	} else {
		left := r.size - r.head - 1
		copy(data, r.data[r.head+1:])
		copy(data[left:], r.data[:toReadLen-left])
	}

	if !peek {
		r.head = headNext
	}

	return toReadLen, nil
}
