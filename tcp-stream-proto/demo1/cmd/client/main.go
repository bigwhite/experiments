package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/bigwhite/tcp-stream-proto/demo1/pkg/frame"
	"github.com/bigwhite/tcp-stream-proto/demo1/pkg/packet"
	"github.com/lucasepe/codename"
)

func main() {
	var wg sync.WaitGroup
	var num int = 5

	wg.Add(5)

	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			startClient()
		}()
	}
	wg.Wait()
}

func startClient() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	rng, err := codename.DefaultRNG()
	if err != nil {
		panic(err)
	}

	frameCodec := frame.NewMyFrameCodec()
	var counter int

	go func() {
		for {
			// handle ack
			// read from the connection
			ackFramePayLoad, err := frameCodec.Decode(conn)
			if err != nil {
				panic(err)
			}

			p, err := packet.Decode(ackFramePayLoad)

			submitAck, ok := p.(*packet.SubmitAck)
			if !ok {
				panic("not submitack")
			}
			fmt.Printf("the result of submit ack[%s] is %d\n", submitAck.ID, submitAck.Result)
		}
	}()

	for {
		// send submit
		counter++
		id := fmt.Sprintf("%08d", counter) // 8 byte string
		payload := codename.Generate(rng, 4)
		s := &packet.Submit{
			ID:      id,
			Payload: []byte(payload),
		}

		framePayload, err := packet.Encode(s)
		if err != nil {
			panic(err)
		}

		fmt.Printf("send submit id = %s, payload=%s, frame length = %d\n", s.ID, s.Payload, len(framePayload)+4)

		err = frameCodec.Encode(conn, framePayload)
		if err != nil {
			panic(err)
		}

		//	time.Sleep(1 * time.Second)
	}
}
