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
var target *string
var wcMu sync.Mutex // sync the access for websocket.Conn

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	signalingAddr = flag.String("signaling-address", "localhost:18080", "address that the signaling server is hosted on.")
	id = flag.String("id", "offer-peer-1", "unique id of the offer peer")
	target = flag.String("target", "", "target id of the other peer")
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
		if cErr := peerConnection.Close(); cErr != nil {
			log.Printf("offer: cannot close peerConnection: %v\n", cErr)
		}
	}()
	log.Println("offer: new peerConnection ok")

	// When an ICE candidate is available send to the other Pion instance
	// the other Pion instance will add this candidate by calling AddICECandidate
	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			return
		}
		log.Printf("offer: invoke peerConnection.OnICECandidate: %#v\n", *c)

		candidatesMux.Lock()
		defer candidatesMux.Unlock()

		desc := peerConnection.RemoteDescription()
		if desc == nil {
			pendingCandidates = append(pendingCandidates, c)
		} else if onICECandidateErr := signalCandidate(wc, *id, *target, c); onICECandidateErr != nil {
			panic(onICECandidateErr)
		}
	})

	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}
	log.Printf("offer: create new channel\n")

	// Set the handler for Peer connection state
	// This will notify you when the peer has connected/disconnected
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("offer: Peer Connection State has changed: %s\n", s.String())

		if s == webrtc.PeerConnectionStateFailed {
			// Wait until PeerConnection has had no network activity for 30 seconds or another failure. It may be reconnected using an ICE Restart.
			// Use webrtc.PeerConnectionStateDisconnected if you are interested in detecting faster timeout.
			// Note that the PeerConnection may come back from PeerConnectionStateDisconnected.
			log.Println("offer: Peer Connection has gone to failed exiting")
			os.Exit(0)
		}

		if s == webrtc.PeerConnectionStateClosed {
			// PeerConnection was explicitly closed. This usually happens from a DTLS CloseNotify
			log.Println("offer: Peer Connection has gone to closed exiting")
			os.Exit(0)
		}
	})

	// Register channel opening handling
	dataChannel.OnOpen(func() {
		log.Printf("offer: Data channel '%s'-'%d' open. Random messages will now be sent to any connected DataChannels every 5 seconds\n", dataChannel.Label(), dataChannel.ID())

		for range time.NewTicker(5 * time.Second).C {
			//message := signal.RandSeq(15)
			message := fmt.Sprintf("offer-%d", rand.Int31()) //signal.RandSeq(15)
			log.Printf("offer: Sending '%s'\n", message)

			// Send the message as text
			sendTextErr := dataChannel.SendText(message)
			if sendTextErr != nil {
				panic(sendTextErr)
			}
		}
	})

	// Register text message handling
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Printf("offer: Message from DataChannel '%s': '%s'\n", dataChannel.Label(), string(msg.Data))
	})

	// communicate with the signaling server
	u := url.URL{Scheme: "ws", Host: *signalingAddr, Path: "/offer"}
	log.Printf("offer: connecting to %s", u.String())

	wc, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("offer: dial %s error: %v", u.String(), err)
	}
	defer wc.Close()

	// Create an offer to send to the other process
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		panic(err)
	}
	log.Printf("offer: create offer\n")

	sdp, err := json.Marshal(offer)
	if err != nil {
		panic(err)
	}

	// send offer to signaling server
	var req = proto.Request{
		SourceID: *id,
		TargetID: *target,
		Body:     sdp,
	}

	var am = proto.Message{
		Cmd: proto.CmdOffer,
	}

	payload, err := req.ToJSON()
	if err != nil {
		log.Fatalf("offer: request payload marshal error: %v", err)
	}
	am.Payload = payload

	message, err := am.ToJSON()
	if err != nil {
		log.Fatalf("offer: message marshal error: %v", err)
	}

	wcMu.Lock()
	err = wc.WriteMessage(websocket.BinaryMessage, message)
	if err != nil {
		wcMu.Unlock()
		log.Fatalln("offer: write message error:", err)
	}
	wcMu.Unlock()

	// offer event loop
	for {
		// read message from signaling
		_, data, err := wc.ReadMessage()
		if err != nil {
			log.Println("offer: read message error:", err)
			return
		}

		var message proto.Message
		err = message.FromJSON(data)
		if err != nil {
			log.Println("offer: message unmarshal error:", err)
			return
		}

		switch message.Cmd {
		case proto.CmdAnswer:
			var req proto.Request
			err = req.FromJSON(message.Payload)
			if err != nil {
				log.Println("offer: unmarshal request error:", err)
				return
			}
			if req.TargetID != *id {
				log.Printf("offer: the target id[%s] of answer request is not me", req.TargetID)
				continue
			}

			log.Printf("offer: recv answer(sdp) message from %s\n", req.SourceID)

			// return response
			err = returnResp(wc, proto.CmdAnswerResp, &proto.Response{
				Code: 0,
				Msg:  "ok",
			})

			// handle answer message
			err = handleAnswer(wc, peerConnection, &req, offer)
			if err != nil {
				log.Println("offer: handle answer message error:", err)
				return
			}
		case proto.CmdCandidate:
			var req proto.Request
			err = req.FromJSON(message.Payload)
			if err != nil {
				log.Println("offer: unmarshal request error:", err)
				return
			}
			if req.TargetID != *id {
				log.Printf("offer: the target id[%s] of candidate request is not me", req.TargetID)
				continue
			}
			log.Printf("offer: recv candidate message from %s\n", req.SourceID)

			// return response
			err = returnResp(wc, proto.CmdCandidateResp, &proto.Response{
				Code: 0,
				Msg:  "ok",
			})

			// handle candidate message
			err = handleCandidate(wc, peerConnection, &req)
			if err != nil {
				log.Println("offer: handle candidate message error:", err)
				return
			}

		case proto.CmdCandidateResp, proto.CmdOfferResp:
			// get response
			resp := proto.Response{}
			err = resp.FromJSON(message.Payload)
			if err != nil {
				log.Fatalln("offer: unmarshal resp message error:", err)
			}

			log.Printf("offer: recv resp[%d]: %#v\n", message.Cmd, resp)
		}
	}

	// Block forever
	select {}
}

func handleAnswer(c *websocket.Conn, peerConnection *webrtc.PeerConnection, req *proto.Request, offer webrtc.SessionDescription) error {
	sdp := webrtc.SessionDescription{}
	if err := json.NewDecoder(bytes.NewReader(req.Body)).Decode(&sdp); err != nil {
		return err
	}

	// Sets the LocalDescription, and starts our UDP listeners
	// Note: this will start the gathering of ICE candidates
	if err := peerConnection.SetLocalDescription(offer); err != nil {
		return err
	}
	log.Printf("offer: set local desc\n")

	if err := peerConnection.SetRemoteDescription(sdp); err != nil {
		return err
	}
	log.Printf("offer: set remote desc\n")

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
