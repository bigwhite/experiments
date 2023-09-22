package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/bigwhite/webrtc/signaling/proto"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:18080", "signaling service address")
var upgrader = websocket.Upgrader{} // use default options

func offer(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil) // *websocket.Conn
	if err != nil {
		log.Print("signaling: websocket upgrade error:", err)
		return
	}
	defer c.Close()

	err = offerPeerEventLoop(c, w)
	if err != nil {
		log.Println("signaling: offerPeerEventLoop error:", err)
		return
	}
	log.Println("signaling: offerPeerEventLoop exit")
}

func offerPeerEventLoop(c *websocket.Conn, w http.ResponseWriter) (err error) {
	var mt int
	var message []byte

	for {
		mt, message, err = c.ReadMessage()
		if err != nil {
			log.Println("signaling: read message from offer peer error:", err)
			return
		}

		// unmarshal offer Message
		var om proto.Message
		err = om.FromJSON(message)
		if err != nil {
			log.Println("signaling: unmarshal message from offer peer error:", err)
			return
		}

		switch om.Cmd {
		case proto.CmdOffer, proto.CmdCandidate:
			// recv a request message from offer peer
			// foward it to answer peer after responsing
			var req proto.Request
			err = req.FromJSON(om.Payload)
			if err != nil {
				log.Println("signaling: unmarshal request from offer peer error:", err)
				return
			}
			log.Printf("signaling: recv request[%d] from offer peer", om.Cmd)

			var rsp proto.Response
			pa, err := pAnswers.findPeer(req.TargetID)
			if err != nil {
				// response to offer peer
				rsp.Code = 1
				rsp.Msg = err.Error()
				returnResp(c, om.Cmd+100, &rsp)
				continue
			}

			rsp.Msg = "ok"
			returnResp(c, om.Cmd+100, &rsp)
			log.Println("signaling: send offer resp ok")

			// store offer peer into pOffers
			p := peerOffer{
				peer: peer{
					id:   req.SourceID,
					conn: c,
				},
				targetID: req.TargetID,
			}
			pOffers.addPeer(&p)
			log.Println("signaling: add offer peer: ", req.SourceID)

			// forward request to answer peer
			err = pa.conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("signaling: forward request to answer peer error:", err)
				return err
			}
			log.Printf("signaling: forward request[%d] to answer peer ok", om.Cmd)
		case proto.CmdAnswerResp:
			log.Println("signaling: recv answer response from offer peer")
		case proto.CmdCandidateResp:
			log.Println("signaling: recv candidate response from offer peer")
		default:
			log.Println("signaling: unsupport cmd:", om.Cmd)
		}
	}
}

func answerPeerEventLoop(c *websocket.Conn, w http.ResponseWriter) (err error) {
	var mt int
	var message []byte

	for {
		mt, message, err = c.ReadMessage()
		if err != nil {
			log.Println("signaling: read message from answer peer error:", err)
			return
		}

		var am proto.Message
		err = am.FromJSON(message)
		if err != nil {
			log.Println("signaling: unmarshal message from answer peer error:", err)
			return
		}

		var req proto.Request
		err = req.FromJSON(am.Payload)
		if err != nil {
			log.Println("signaling: unmarshal request from answer peer error:", err)
			return
		}

		var rsp proto.Response
		var rspData []byte

		switch am.Cmd {
		case proto.CmdInit:
			// store the answer peer into answerMap
			p := peerAnswer{
				peer: peer{
					id:   req.SourceID,
					conn: c,
				},
			}
			rsp.Code = 0
			rsp.Msg = "ok"
			pAnswers.addPeer(&p)
			log.Println("signaling: add answer peer:", p.id)
			rspData, err = rsp.ToJSON()
			if err != nil {
				log.Println("signaling: marshal init response error:", err)
				return
			}

			var message = proto.Message{
				Cmd:     proto.CmdInitResp,
				Payload: rspData,
			}
			var messageData []byte
			messageData, err = message.ToJSON()
			if err != nil {
				log.Println("signaling: marshal init response message error:", err)
				return
			}

			err = c.WriteMessage(mt, messageData)
			if err != nil {
				log.Println("signaling: write response message to answer peer error:", err)
				return
			}

		case proto.CmdAnswer, proto.CmdCandidate:
			log.Printf("signaling: recv request[%d] from answer peer", am.Cmd)
			// recv answer or candidate from answer peer
			// foward to offer peer after responsing
			po, err := pOffers.findPeer(req.TargetID)
			if err != nil {
				log.Printf("signaling: not find the target offer[%s]", req.TargetID)
				rsp.Code = 1
				rsp.Msg = err.Error()
				returnResp(c, am.Cmd+100, &rsp)
				continue
			}
			rsp.Msg = "ok"
			rsp.Code = 0
			returnResp(c, am.Cmd+100, &rsp)

			// forward message to offer peer
			err = po.conn.WriteMessage(mt, message)
			if err != nil {
				log.Printf("signaling: forward request[%d] to offer peer error: %v", am.Cmd, err)
				return err
			}
			log.Printf("signaling: forward request[%d] to offer peer[%s] ok", am.Cmd, po.id)
		case proto.CmdOfferResp:
			log.Println("signaling: recv offer response from answer peer")
		case proto.CmdCandidateResp:
			log.Println("signaling: recv candidate response from answer peer")
		default:
			log.Println("signaling: unsupport cmd:", am.Cmd)
		}
	}
}

// in a standalone goroutine
func register(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil) // *websocket.Conn
	if err != nil {
		log.Print("signaling: websocket upgrade error:", err)
		return
	}
	defer c.Close()

	err = answerPeerEventLoop(c, w)
	if err != nil {
		log.Println("signaling: answerPeerEventLoop error:", err)
		return
	}
	log.Println("signaling: answerPeerEventLoop exit")
}

type peer struct {
	id   string // unique id for identifying peer
	conn *websocket.Conn
}

type peerOffer struct {
	peer
	targetID string // the peer id which peerOffer want to connect
}

type peerAnswer struct {
	peer
}

type peerOffers struct {
	m map[string]*peerOffer
	sync.Mutex
}

var pOffers = peerOffers{
	m: make(map[string]*peerOffer),
}

func (prs *peerOffers) addPeer(p *peerOffer) {
	prs.Lock()
	defer prs.Unlock()
	prs.m[p.id] = p
}

func (prs *peerOffers) findPeer(id string) (*peerOffer, error) {
	prs.Lock()
	defer prs.Unlock()
	p, ok := prs.m[id]
	if ok {
		return p, nil
	}

	return nil, os.ErrNotExist
}

var pAnswers = peerAnswers{
	m: make(map[string]*peerAnswer),
}

type peerAnswers struct {
	m map[string]*peerAnswer
	sync.Mutex
}

func (prs *peerAnswers) addPeer(p *peerAnswer) {
	prs.Lock()
	defer prs.Unlock()
	prs.m[p.id] = p
}

func (prs *peerAnswers) findPeer(id string) (*peerAnswer, error) {
	prs.Lock()
	defer prs.Unlock()
	p, ok := prs.m[id]
	if ok {
		return p, nil
	}

	return nil, os.ErrNotExist
}

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	http.HandleFunc("/register", register) // for peerAnswer
	http.HandleFunc("/offer", offer)       // for peerOffer
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func returnResp(c *websocket.Conn, cmd int, resp *proto.Response) error {
	b, err := resp.ToJSON()
	if err != nil {
		return err
	}

	var message = proto.Message{
		Cmd:     cmd,
		Payload: b,
	}

	data, err := message.ToJSON()
	if err != nil {
		return err
	}

	return c.WriteMessage(websocket.BinaryMessage, data)
}
