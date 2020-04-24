package protobufbench

import (
	"fmt"
	"os"
	"testing"

	"github.com/bigwhite/protobufbench_gogoprotofaster/submit"
	"github.com/gogo/protobuf/proto"
)

var request = submit.Request{
	Recvtime: 170123456,
	Uniqueid: "a1b2c3d4e5f6g7h8i9",
	Token:    "xxxx-1111-yyyy-2222-zzzz-3333",
	Phone:    "13900010002",
	Content:  "Customizing the fields of the messages to be the fields that you actually want to use removes the need to copy between the structs you use and structs you use to serialize. gogoprotobuf also offers more serialization formats and generation of tests and even more methods.",
	Sign:     "tonybaiXZYDFDS",
	Type:     "submit",
	Extend:   "",
	Version:  "v1.0.0",
}

var requestToUnMarshal []byte

func init() {
	var err error
	requestToUnMarshal, err = proto.Marshal(&request)
	if err != nil {
		fmt.Printf("marshal err:%s\n", err)
		os.Exit(1)
	}
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(&request)
	}
}
func BenchmarkUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var request submit.Request
	for i := 0; i < b.N; i++ {
		_ = proto.Unmarshal(requestToUnMarshal, &request)

	}
}

func BenchmarkMarshalInParalell(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = proto.Marshal(&request)
		}
	})
}
func BenchmarkUnmarshalParalell(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var request submit.Request
			_ = proto.Unmarshal(requestToUnMarshal, &request)
		}
	})
}
