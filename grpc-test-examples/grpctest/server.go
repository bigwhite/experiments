package main

import (
	"context"
	"fmt"
	"strconv"

	pb "demo/mygrpc"
)

type server struct{}

func (s *server) UnaryRPC(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {
	message := "Unary RPC received: " + req.Message
	fmt.Println(message)

	return &pb.ResponseMessage{
		Message: "Unary RPC response",
	}, nil
}

func (s *server) ServerStreamingRPC(req *pb.RequestMessage, stream pb.MyService_ServerStreamingRPCServer) error {
	message := "Server Streaming RPC received: " + req.Message
	fmt.Println(message)

	for i := 0; i < 5; i++ {
		response := &pb.ResponseMessage{
			Message: "Server Streaming RPC response " + strconv.Itoa(i+1),
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) ClientStreamingRPC(stream pb.MyService_ClientStreamingRPCServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		messages = append(messages, req.Message)

		if req.Message == "end" {
			break
		}
	}

	message := "Client Streaming RPC received: " + fmt.Sprintf("%v", messages)
	fmt.Println(message)

	return stream.SendAndClose(&pb.ResponseMessage{
		Message: "Client Streaming RPC response",
	})
}

func (s *server) BidirectionalStreamingRPC(stream pb.MyService_BidirectionalStreamingRPCServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		message := "Bidirectional Streaming RPC received: " + req.Message
		fmt.Println(message)

		response := &pb.ResponseMessage{
			Message: "Bidirectional Streaming RPC response",
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
