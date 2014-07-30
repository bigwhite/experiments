/*
 * Copyright 2012 iWobi.net. All rights reserved.
 *
 * proto parser
 *
 *
 * SPECIFICATIONS
 *    request:
 *      incr key value\r\n
 *
 *    reply:
 *      - timestamp NOT_FOUND\r\n to indicate the item with this value was not found
 *      - timestamp <value>\r\n
 *
 *    note: timestamp is the local time of host which counterproxy is deployed on,
 *          its format is "YYYYmmddhhMMSS"
 */

package proto

import (
	"errors"
	"iwobi.net/utils/ringbuf"
	"strings"
)

const (
	INCR = "incr"
)

type ProtoPackage struct {
	Cmd string
	Key string
	Val string
}

var ErrInvalidPack = errors.New("protocol package is invalid")
var ErrInComplete = errors.New("data in buffer is incomplete")

func Parse(rb *ringbuf.RingBuf) (packages []ProtoPackage, err error) {
	var data [128]byte

	if rb.Len() == 0 {
		return packages, nil
	}

	for {
		n, _ := rb.Read(data[:], len(data), true)
		if n == 0 {
			break
		}

		i := strings.Index(string(data[:n]), "\r\n")
		if i == -1 {
			return packages, ErrInComplete
		} else {
			s := strings.Split(string(data[:i]), " ")
			if len(s) != 3 || s[0] != INCR {
				return packages, ErrInvalidPack
			}

			var pp ProtoPackage
			pp.Cmd, pp.Key, pp.Val = s[0], s[1], s[2]
			packages = append(packages, pp)
			rb.Read(data[:], i+2, false)
		}
	}
	return packages, nil
}
