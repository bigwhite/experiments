package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/minio/simdjson-go"

	"github.com/bigwhite/codec/proto"
	pb "github.com/gogo/protobuf/proto"
)

var jsonText = []byte(`{
  "clientid" : "xxxxyyyyzzzz123456789123456789",
  "topic" : "vsp/vehicles/notify",
  "data" : "5ZGo5oGp5p2l77yIMTg5OOW5tDPmnIg15pelLTE5NzblubQx5pyIOOaXpe+8iSDvvIzlrZfnv5TlrofvvIzmm77nlKjlkI3po57po57jgIHkvI3osarjgIHlsJHlsbHjgIHlhqDnlJ/nrYkgWzFdICBbNl0gIO+8jOWOn+exjea1meaxn+e7jeWFtO+8jDE4OTjlubQz5pyINeaXpeeUn+S6juaxn+iLj+a3ruWuieOAgjE5MjHlubTliqDlhaXkuK3lm73lhbHkuqflhZrvvIzmmK/kvJ/lpKfnmoTpqazlhYvmgJ3kuLvkuYnogIXvvIzkvJ/lpKfnmoTml6DkuqfpmLbnuqfpnanlkb3lrrbjgIHmlL/msrvlrrbjgIHlhpvkuovlrrbjgIHlpJbkuqTlrrbvvIzlhZrlkozlm73lrrbkuLvopoHpooblr7zkurrkuYvkuIDvvIzkuK3lm73kurrmsJHop6PmlL7lhpvkuLvopoHliJvlu7rkurrkuYvkuIDvvIzkuK3ljY7kurrmsJHlhbHlkozlm73nmoTlvIDlm73lhYPli4vvvIzmmK/ku6Xmr5vms73kuJzlkIzlv5fkuLrmoLjlv4PnmoTlhZrnmoTnrKzkuIDku6PkuK3lpK7pooblr7zpm4bkvZPnmoTph43opoHmiJDlkZggWzJdICDjgIIKMTk3NuW5tDHmnIg45pel5Zyo5YyX5Lqs6YCd5LiW44CC5LuW55qE6YCd5LiW5Y+X5Yiw5p6B5bm/5rOb55qE5oK85b+144CC55Sx5LqO5LuW5LiA6LSv5Yuk5aWL5bel5L2c77yM5Lil5LqO5b6L5bex77yM5YWz5b+D576k5LyX77yM6KKr56ew5Li64oCc5Lq65rCR55qE5aW95oC755CG4oCd44CC5LuW55qE5Li76KaB6JGX5L2c5pS25YWl44CK5ZGo5oGp5p2l6YCJ6ZuG44CL44CC"}`)

type JsonMessage struct {
	ClientId string `json:"clientid"`
	Topic    string `json:"topic"`
	Payload  []byte `json:"data"`
}

func simdjsonUnmarshal(text []byte, m *JsonMessage) error {
	/*
		if !simdjson.SupportedCPU() {
			log.Fatal("Unsupported CPU")
		}*/

	parsed, err := simdjson.Parse(text, nil)
	if err != nil {
		return err
	}
	_ = parsed
	return nil
}

func BenchmarkSimdJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	//var m JsonMessage
	for i := 0; i < b.N; i++ {
		//_ = simdjsonUnmarshal(jsonText, &m)
		simdjson.Parse(jsonText, nil)
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var m JsonMessage
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(jsonText, &m)
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	b.ReportAllocs()
	var m = JsonMessage{
		ClientId: "xxxxyyyyzzzz123456789123456789",
		Topic:    "vsp/vehicles/notify",
		Payload:  jsonText,
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(&m)
	}
}

func BenchmarkSonicJsonUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var m JsonMessage
	for i := 0; i < b.N; i++ {
		sonic.Unmarshal(jsonText, &m)
	}
}

func BenchmarkSonicJsonMarshal(b *testing.B) {
	b.ReportAllocs()
	var m = JsonMessage{
		ClientId: "xxxxyyyyzzzz123456789123456789",
		Topic:    "vsp/vehicles/notify",
		Payload:  jsonText,
	}
	for i := 0; i < b.N; i++ {
		sonic.Marshal(&m)
	}
}

var msgForPb = proto.Message{
	Clientid: "xxxxyyyyzzzz123456789123456789",
	Topic:    "vsp/vehicles/notify",
	Payload:  string(jsonText),
}

var data []byte

func init() {
	data, _ = pb.Marshal(&msgForPb)
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	b.ReportAllocs()
	var m proto.Message

	for i := 0; i < b.N; i++ {
		_ = pb.Unmarshal(data, &m)
	}
}

func BenchmarkProtobufMarshal(b *testing.B) {
	b.ReportAllocs()
	var m = proto.Message{
		Clientid: "xxxxyyyyzzzz123456789123456789",
		Topic:    "vsp/vehicles/notify",
		Payload:  string(jsonText),
	}

	for i := 0; i < b.N; i++ {
		pb.Marshal(&m)
	}
}
