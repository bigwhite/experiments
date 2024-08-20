package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/logging"
	"github.com/pion/webrtc/v3"
)

type signalMsg struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

var (
	signalingServer string
	roomID          string
)

func init() {
	flag.StringVar(&signalingServer, "server", "ws://localhost:28080/ws", "Signaling server WebSocket URL")
	flag.StringVar(&roomID, "room", "", "Room ID (leave empty to create a new room)")
	flag.Parse()
}

func main() {
	// Connect to signaling server
	signalingURL := fmt.Sprintf("%s?room=%s", signalingServer, roomID)
	conn, _, err := websocket.DefaultDialer.Dial(signalingURL, nil)
	if err != nil {
		log.Fatal("Error connecting to signaling server:", err)
	}
	defer conn.Close()
	log.Println("connect to signaling server ok")

	// Create a new RTCPeerConnection
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// 创建一个自定义的日志工厂
	loggerFactory := logging.NewDefaultLoggerFactory()
	loggerFactory.DefaultLogLevel = logging.LogLevelTrace
	//loggerFactory.DefaultLogLevel = logging.LogLevelInfo
	//loggerFactory.DefaultLogLevel = logging.LogLevelDebug

	// Enable detailed logging
	s := webrtc.SettingEngine{}
	s.LoggerFactory = loggerFactory
	s.SetICETimeouts(5*time.Second, 5*time.Second, 5*time.Second)

	api := webrtc.NewAPI(webrtc.WithSettingEngine(s))
	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create a datachannel
	dataChannel, err := peerConnection.CreateDataChannel("test", nil)
	if err != nil {
		log.Fatal(err)
	}

	dataChannel.OnOpen(func() {
		log.Println("Data channel is open")
		go func() {
			for {
				err := dataChannel.SendText("Hello from " + roomID)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(5 * time.Second)
			}
		}()
	})

	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Printf("Received message: %s\n", string(msg.Data))
	})

	// Set the handler for ICE connection state
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Printf("ICE Connection State has changed: %s\n", connectionState.String())
	})

	// Set the handler for Peer connection state
	peerConnection.OnConnectionStateChange(func(s webrtc.PeerConnectionState) {
		log.Printf("Peer Connection State has changed: %s\n", s.String())
	})

	// Set the handler for Signaling state
	peerConnection.OnSignalingStateChange(func(s webrtc.SignalingState) {
		log.Printf("Signaling State has changed: %s\n", s.String())
	})

	// Register data channel creation handling
	peerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
		log.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		d.OnOpen(func() {
			log.Printf("Data channel '%s'-'%d' open.\n", d.Label(), d.ID())
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			log.Printf("Message from DataChannel '%s': '%s'\n", d.Label(), string(msg.Data))
		})
	})

	// Set the handler for ICE candidate generation
	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		if i == nil {
			return
		}

		candidateString, err := json.Marshal(i.ToJSON())
		if err != nil {
			log.Println(err)
			return
		}

		if writeErr := conn.WriteJSON(&signalMsg{
			Type: "candidate",
			Data: string(candidateString),
		}); writeErr != nil {
			log.Println(writeErr)
		}
	})

	// Handle incoming messages from signaling server
	go func() {
		for {
			_, rawMsg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}
			log.Println("recv msg from signaling server")

			var msg signalMsg
			if err := json.Unmarshal(rawMsg, &msg); err != nil {
				log.Println("Error parsing message:", err)
				continue
			}
			log.Println("recv msg is", msg)

			switch msg.Type {
			case "offer":
				log.Println("recv a offer msg")
				offer := webrtc.SessionDescription{}
				if err := json.Unmarshal([]byte(msg.Data), &offer); err != nil {
					log.Println("Error parsing offer:", err)
					continue
				}

				if err := peerConnection.SetRemoteDescription(offer); err != nil {
					log.Println("Error setting remote description:", err)
					continue
				}

				answer, err := peerConnection.CreateAnswer(nil)
				if err != nil {
					log.Println("Error creating answer:", err)
					continue
				}

				if err := peerConnection.SetLocalDescription(answer); err != nil {
					log.Println("Error setting local description:", err)
					continue
				}

				answerString, err := json.Marshal(answer)
				if err != nil {
					log.Println("Error encoding answer:", err)
					continue
				}

				if err := conn.WriteJSON(&signalMsg{
					Type: "answer",
					Data: string(answerString),
				}); err != nil {
					log.Println("Error sending answer:", err)
				}
				log.Println("send answer ok")

			case "answer":
				log.Println("recv a answer msg")
				answer := webrtc.SessionDescription{}
				if err := json.Unmarshal([]byte(msg.Data), &answer); err != nil {
					log.Println("Error parsing answer:", err)
					continue
				}

				if err := peerConnection.SetRemoteDescription(answer); err != nil {
					log.Println("Error setting remote description:", err)
				}
				log.Println("set remote desc for answer ok")

			case "candidate":
				candidate := webrtc.ICECandidateInit{}
				if err := json.Unmarshal([]byte(msg.Data), &candidate); err != nil {
					log.Println("Error parsing candidate:", err)
					continue
				}

				if err := peerConnection.AddICECandidate(candidate); err != nil {
					log.Println("Error adding ICE candidate:", err)
				}
				log.Println("adding ICE candidate:", candidate)
			}
		}
	}()

	// Create an offer if we are the peer to join the room
	if roomID != "" {
		offer, err := peerConnection.CreateOffer(nil)
		if err != nil {
			log.Fatal(err)
		}

		if err := peerConnection.SetLocalDescription(offer); err != nil {
			log.Fatal(err)
		}

		offerString, err := json.Marshal(offer)
		if err != nil {
			log.Fatal(err)
		}

		if err := conn.WriteJSON(&signalMsg{
			Type: "offer",
			Data: string(offerString),
		}); err != nil {
			log.Fatal(err)
		}
		log.Printf("send offer to signaling server ok\n")
	}

	// Wait forever
	select {}
}
