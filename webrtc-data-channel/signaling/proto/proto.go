package proto

import "encoding/json"

const (
	// originated from answer peer
	CmdInit = iota + 1
	CmdAnswer

	// originated from answer peer
	CmdOffer

	// from both peer
	CmdCandidate
)

const (
	CmdInitResp = iota + 101 // CmdInit + 100
	CmdAnswerResp
	CmdOfferResp
	CmdCandidateResp
)

type Message struct {
	Cmd     int    `json:"command"`
	Payload []byte `json:"payload"` // carry all kinds of request and response
}

func (m Message) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

func (m *Message) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// Request is one kind of payload for Message
type Request struct {
	SourceID string `json:"source"`
	TargetID string `json:"target"`
	Body     []byte `json:"body"` // carry register, offer, answer, candidate
}

func (m Request) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

func (m *Request) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// Request is another payload for Message
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (m Response) ToJSON() ([]byte, error) {
	return json.Marshal(&m)
}

func (m *Response) FromJSON(data []byte) error {
	return json.Unmarshal(data, m)
}
