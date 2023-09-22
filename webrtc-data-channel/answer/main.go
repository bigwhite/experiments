package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/bigwhite/webrtc/signaling/proto"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

func signalCandidate(wc *websocket.Conn, source, target string, c *webrtc.ICECandidate) error {
	payload := []byte(c.ToJSON().Candidate)

	var req = proto.Request{
		SourceID: source,
		TargetID: target,
		Body:     payload,
	}

	reqData, err := req.ToJSON()
	if err != nil {
		return err
	}

	var message = proto.Message{
		Cmd:     proto.CmdCandidate,
		Payload: reqData,
	}

	messageData, err := message.ToJSON()
	if err != nil {
		return err
	}

	wcMu.Lock()
	defer wcMu.Unlock()
	return wc.WriteMessage(websocket.BinaryMessage, messageData)
}

var signalingAddr *string
var id *string
var target string
var wcMu sync.Mutex // sync the access for websocket.Conn

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	signalingAddr = flag.String("signaling-address", "localhost:18080", "address that the signaling server is hosted on.")
	id = flag.String("id", "answer-peer-1", "unique id of the answer peer")
	flag.Parse()

	var wc *websocket.Conn
	var err error

	var candidatesMux sync.Mutex
	pendingCandidates := make([]*webrtc.ICECandidate, 0)

	// Prepare the configuration
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				//URLs: []string{"stun:stun.l.google.com:19302"},
				URLs: []string{"stun:74.125.137.127:19302"},
			},
		},
	}

	// Create a new RTCPeerConnection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Printf("answer: cannot close peerConnection: %v\n", err)
		}
	}()
	log.Println("answer: NewPeerConnection ok")

	// When an ICE candidate is available send to the other Pion instance
	// the other Pion instance will add this candidate by calling AddICECandidate
	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}
		log.Printf("answer: invoke peerConnection.OnICECandidate: %s\n", c.Address)

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		desc := peerConnection.RemoteDescription()
		if desc == nil {
			pendingCandidates = append(pendingCandidates, c)
		} else if onICECandidateErr := signalCandidate(wc, *id, target, c); onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	})

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("answer: Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			log.Println("answer: Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if s == webrtc.PeerConnectionStateClosed {
			// PeerConnection was explicitly closed. This usually happens from a DTLS CloseNotify
			log.Println("answer: Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Printf("answer: New DataChannel %s %d\n", d.Label(), d.ID())

		// Register channel opening handling
		d.OnOpen(func() {
			log.Printf("answer: Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", d.Label(), d.ID())

			for range time.NewTicker(5 * time.Second).C {
				message := fmt.Sprintf("answer-%d", rand.Int31()) //signal.RandSeq(15)
				log.Printf("answer: Sending '%s'\n", message)

				// Send the message as text
				sendTextErr := d.SendText(message)
				if sendTextErr != nil {
					panic(sendTextErr)
				}
			}
		})

		// Register text message handling
		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			log.Printf("answer: message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// communicate with the signaling server
	u := url.URL{Scheme: "ws", Host: *signalingAddr, Path: "/register"}
	log.Printf("answer: connecting to %s", u.String())

	wc, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("answer: dial %s error: %v", u.String(), err)
	}
	defer wc.Close()

	// send register message
	var req = proto.Request{
		SourceID: *id,
	}

	var am = proto.Message{
		Cmd: proto.CmdInit,
	}

	payload, err := req.ToJSON()
	if err != nil {
		log.Fatalf("answer: request payload marshal error: %v", err)
	}
	am.Payload = payload

	message, err := am.ToJSON()
	if err != nil {
		log.Fatalf("answer: message marshal error: %v", err)
	}

	wcMu.Lock()
	err = wc.WriteMessage(websocket.BinaryMessage, message)
	if err != nil {
		wcMu.Unlock()
		log.Fatalln("answer: write message error:", err)
	}
	wcMu.Unlock()

	// answer event loop
	for {
		// read message from signaling
		_, data, err := wc.ReadMessage()
		if err != nil {
			log.Println("answer: read message error:", err)
			return
		}

		var message proto.Message
		err = message.FromJSON(data)
		if err != nil {
			log.Println("answer: message unmarshal error:", err)
			return
		}

		var req proto.Request
		err = req.FromJSON(message.Payload)
		if err != nil {
			log.Println("answer: unmarshal request error:", err)
			return
		}

		switch message.Cmd {
		case proto.CmdOffer:
			if req.TargetID != *id {
				log.Printf("answer: the target id[%s] of offer request is not me", req.TargetID)
				continue
			}

			log.Printf("answer: recv offer message from %s\n", req.SourceID)
			target = req.SourceID

			// return response
			err = returnResp(wc, proto.CmdOfferResp, &proto.Response{
				Code: 0,
				Msg:  "ok",
			})

			// handle offer message
			err = handleOffer(wc, peerConnection, &req)
			if err != nil {
				log.Println("answer: handle offer message error:", err)
				return
			}
		case proto.CmdCandidate:
			if req.TargetID != *id {
				log.Printf("answer: the target id[%s] of candidate request is not me", req.TargetID)
				continue
			}
			log.Printf("answer: recv candidate message from %s\n", req.SourceID)

			// return response
			err = returnResp(wc, proto.CmdCandidateResp, &proto.Response{
				Code: 0,
				Msg:  "ok",
			})

			// handle candidate message
			err = handleCandidate(wc, peerConnection, &req)
			if err != nil {
				log.Println("answer: handle candidate message error:", err)
				return
			}

		case proto.CmdInitResp, proto.CmdCandidateResp, proto.CmdAnswerResp:
			// get response
			resp := proto.Response{}
			err = resp.FromJSON(message.Payload)
			if err != nil {
				log.Fatalln("answer: unmarshal resp message error:", err)
			}

			log.Printf("answer: recv resp[%d]: %#v\n", message.Cmd, resp)
		}
	}
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

	wcMu.Lock()
	defer wcMu.Unlock()
	return c.WriteMessage(websocket.BinaryMessage, data)
}

func handleOffer(c *websocket.Conn, peerConnection *webrtc.PeerConnection, req *proto.Request) error {
	sdp := webrtc.SessionDescription{}
	if err := json.NewDecoder(bytes.NewReader(req.Body)).Decode(&sdp); err != nil {
		return err
	}

	if err := peerConnection.SetRemoteDescription(sdp); err != nil {
		log.Println("answer: set remote description error:", err)
		return err
	}

	// Create an answer to send to the other process
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		log.Println("answer: create answer error:", err)
		return err
	}

	// Send our answer to the signaling server
	payload, err := json.Marshal(answer)
	if err != nil {
		return err
	}

	var answerReq = proto.Request{
		SourceID: req.TargetID,
		TargetID: req.SourceID,
		Body:     payload,
	}

	answerReqData, err := answerReq.ToJSON()
	if err != nil {
		return err
	}

	var message = proto.Message{
		Cmd:     proto.CmdAnswer,
		Payload: answerReqData,
	}

	data, err := message.ToJSON()
	if err != nil {
		return err
	}

	wcMu.Lock()
	err = c.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		wcMu.Unlock()
		log.Println("answer: write answer error: ", err)
		return err
	}
	wcMu.Unlock()

	log.Println("answer: send sdp answer")

	go func() {
		//time.Sleep(5 * time.Second)
		// Sets the LocalDescription, and starts our UDP listeners
		//
		// trigger communication with ice
		err = peerConnection.SetLocalDescription(answer)
		log.Println("answer: set local desc")
		if err != nil {
			log.Println("answer: set local desc error:", err)
			//return err
		}
	}()

	return nil
}

func handleCandidate(c *websocket.Conn, peerConnection *webrtc.PeerConnection, req *proto.Request) error {
	candidate, candidateErr := io.ReadAll(bytes.NewReader(req.Body))
	if candidateErr != nil {
		panic(candidateErr)
	}
	if candidateErr := peerConnection.AddICECandidate(webrtc.ICECandidateInit{Candidate: string(candidate)}); candidateErr != nil {
		return candidateErr
	}

	return nil
}
