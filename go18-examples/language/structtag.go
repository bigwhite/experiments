package main

import "fmt"

type XmlEventRegRequest struct {
	AppID     string `xml:"appid"`
	NeedReply int    `xml:"Reply,omitempty"`
}

type JsonEventRegRequest struct {
	AppID     string `json:"appid"`
	NeedReply int    `json:"reply,omitempty"`
}

func convert(in *XmlEventRegRequest) *JsonEventRegRequest {
	out := &JsonEventRegRequest{}
	*out = (JsonEventRegRequest)(*in)
	return out
}

func main() {
	in := XmlEventRegRequest{
		AppID:     "wx12345678",
		NeedReply: 1,
	}
	out := convert(&in)
	fmt.Println(out)
}
