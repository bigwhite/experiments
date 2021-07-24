package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/panjf2000/gnet"
)

type Frame []byte

// Encode ...
func (cc Frame) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	fmt.Println("into frame encode")
	fmt.Println("in frame encode, buf length is", len(buf))
	fmt.Println("in frame encode, buf is", buf)
	result := make([]byte, 0)

	buffer := bytes.NewBuffer(result)

	// take out the param
	item := c.Context().(Frame)
	fmt.Println("in frame encode, item length is", len(item))
	fmt.Println("in frame encode, item is", item)

	length := uint32(4 + len([]byte(item)))
	if err := binary.Write(buffer, binary.BigEndian, length); err != nil {
		s := fmt.Sprintf("Pack length error , %v", err)
		return nil, errors.New(s)
	}

	n, err := buffer.Write(item)
	if err != nil {
		s := fmt.Sprintf("Pack frame body error , %v", err)
		return nil, errors.New(s)
	}

	if n != len(item) {
		s := fmt.Sprintf("Pack frame body length error , %v", err)
		return nil, errors.New(s)
	}

	return buffer.Bytes(), nil
}

// Decode ...
func (cc Frame) Decode(c gnet.Conn) ([]byte, error) {
	fmt.Println("into frame decode") //只是用来预读取，检查frame完整性
	// read length
	var length uint32
	if size, header := c.ReadN(4); size == 4 {
		byteBuffer := bytes.NewBuffer(header)
		_ = binary.Read(byteBuffer, binary.BigEndian, &length)
		fmt.Println("in frame decode: frame length =", length)

		if size1, body := c.ReadN(int(length)); size1 == int(length) {
			fmt.Println("in frame decode: body length =", len(body))
			c.ShiftN(4 + int(length)) // shift frame header length
			return body, nil
		} else {
			return nil, errors.New("not enough body data")
		}
	}
	return nil, errors.New("not enough length data")
}
