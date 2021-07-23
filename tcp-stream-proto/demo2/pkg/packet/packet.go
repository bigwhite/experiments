package packet

import (
	"bytes"
	"fmt"
)

// Packet协议定义

/*
### packet header
1 byte: commandID

### submit packet

8字节 ID 字符串
任意字节 payload

### submit ack packet

8字节 ID 字符串
1字节 result
*/

const (
	CommandConn   = iota + 0x01 // 0x01
	CommandSubmit               // 0x02
)

const (
	CommandConnAck   = iota + 0x80 // 0x81
	CommandSubmitAck               //0x82
)

type Packet interface {
	Decode([]byte) error     // []byte -> struct
	Encode() ([]byte, error) //  struct -> []byte
}

type PktHdr struct {
	CommandID uint8
}

type Submit struct {
	ID      string
	Payload []byte
}

func (s *Submit) Decode(packet []byte) error {
	s.ID = string(packet[:8])
	s.Payload = packet[8:]
	return nil
}

func (s *Submit) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), s.Payload}, nil), nil
}

type SubmitAck struct {
	ID     string
	Result uint8
}

func (s *SubmitAck) Decode(packet []byte) error {
	s.ID = string(packet[0:8])
	s.Result = uint8(packet[8])
	return nil
}

func (s *SubmitAck) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(s.ID[:8]), []byte{s.Result}}, nil), nil
}

func Decode(packet []byte) (Packet, error) {
	commandID := packet[0]

	switch commandID {
	case CommandConn:
		return nil, nil
	case CommandConnAck:
		return nil, nil
	case CommandSubmit:
		s := Submit{}
		err := s.Decode(packet[1:])
		if err != nil {
			return nil, err
		}
		return &s, nil
	case CommandSubmitAck:
		s := SubmitAck{}
		err := s.Decode(packet[1:])
		if err != nil {
			return nil, err
		}
		return &s, nil
	default:
		return nil, fmt.Errorf("unknown commandID [%d]", commandID)
	}
}

func Encode(p Packet) ([]byte, error) {
	var commandID uint8
	var body []byte
	var err error

	switch t := p.(type) {
	case *Submit:
		commandID = CommandSubmit
		body, err = p.Encode()
		if err != nil {
			return nil, err
		}
	case *SubmitAck:
		commandID = CommandSubmitAck
		body, err = p.Encode()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown type [%s]", t)
	}
	return bytes.Join([][]byte{[]byte{commandID}, body}, nil), nil
}
