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
func (cc Frame) Encode(c gnet.Conn, framePayload []byte) ([]byte, error) {
	result := make([]byte, 0)

	buffer := bytes.NewBuffer(result)

	// encode frame length(4+ framePayload length)
	length := uint32(4 + len([]byte(framePayload)))
	if err := binary.Write(buffer, binary.BigEndian, length); err != nil {
		s := fmt.Sprintf("Pack length error , %v", err)
		return nil, errors.New(s)
	}

	// encode frame payload
	n, err := buffer.Write(framePayload)
	if err != nil {
		s := fmt.Sprintf("Pack frame payload error , %v", err)
		return nil, errors.New(s)
	}

	if n != len(framePayload) {
		s := fmt.Sprintf("Pack frame payload length error , %v", err)
		return nil, errors.New(s)
	}

	return buffer.Bytes(), nil
}

// Decode ...
func (cc Frame) Decode(c gnet.Conn) ([]byte, error) {
	//只是用来预读取，检查frame完整性
	// read length
	var frameLength uint32
	if n, header := c.ReadN(4); n == 4 {
		byteBuffer := bytes.NewBuffer(header)
		_ = binary.Read(byteBuffer, binary.BigEndian, &frameLength)

		if frameLength > 100 {
			c.ResetBuffer()
			return nil, errors.New("length value is wrong")
		}

		if n, wholeFrame := c.ReadN(int(frameLength)); n == int(frameLength) {
			c.ShiftN(int(frameLength)) // shift frame length
			return wholeFrame[4:], nil // return frame payload
		} else {
			return nil, errors.New("not enough frame payload data")
		}
	}
	return nil, errors.New("not enough frame length data")
}
