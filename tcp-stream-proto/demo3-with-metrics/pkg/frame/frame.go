package frame

import (
	"encoding/binary"
	"io"

	"github.com/bytedance/gopkg/lang/mcache"
)

/*
Frame定义

frameHeader + framePayload(packet)

frameHeader
	4 bytes: length 整型，帧总长度(含头及payload)

framePayload
	Packet
*/

type FramePayload []byte

type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error   // data -> frame，并写入io.Writer
	Decode(io.Reader) (FramePayload, error) // 从io.Reader中提取frame payload，并返回给上层
}

type myFrameCodec struct{}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, payload FramePayload) error {
	var f = payload
	var length int32 = int32(len(payload)) + 4

	err := binary.Write(w, binary.BigEndian, &length)
	if err != nil {
		return err
	}

	// make sure all data will be written to outbound stream
	for {
		n, err := w.Write([]byte(f)) // write the frame payload to outbound stream
		if err != nil {
			return err
		}
		if n >= len(f) {
			break
		}
		if n < len(f) {
			f = f[n:]
		}
	}
	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	buf := mcache.Malloc(int(totalLen - 4))
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	return FramePayload(buf), nil
}
