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
	//item := c.Context().(Frame)
	item := buf
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
	var frameLength uint32
	if n, header := c.ReadN(4); n == 4 {
		byteBuffer := bytes.NewBuffer(header)
		_ = binary.Read(byteBuffer, binary.BigEndian, &frameLength)
		fmt.Println("in frame decode: frame length =", frameLength)

		if frameLength > 100 {
			fmt.Printf("in frame decode: frame length[%d] is wrong\n", frameLength)
			c.ResetBuffer()
			return nil, errors.New("length value is wrong")
		}

		if n, wholeFrame := c.ReadN(int(frameLength)); n == int(frameLength) {
			fmt.Println("in frame decode: payload length =", len(wholeFrame)-4)
			c.ShiftN(int(frameLength)) // shift frame header length
			return wholeFrame[4:], nil
		} else {
			return nil, errors.New("not enough payload data")
		}
	}
	return nil, errors.New("not enough frame length data")
}
