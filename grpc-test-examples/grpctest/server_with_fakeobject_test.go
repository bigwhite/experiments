package main

import (
	"testing"

	pb "demo/mygrpc"

	"google.golang.org/grpc"
)

// Fake implementation of ServerStreamingRPC stream
type fakeServerStreamingRPCStream struct {
	grpc.ServerStream
	responses []*pb.ResponseMessage
}

func (m *fakeServerStreamingRPCStream) Send(resp *pb.ResponseMessage) error {
	m.responses = append(m.responses, resp)
	return nil
}

func TestServerServerStreamingRPC(t *testing.T) {
	s := &server{}

	req := &pb.RequestMessage{
		Message: "Test message",
	}

	stream := &fakeServerStreamingRPCStream{}

	err := s.ServerStreamingRPC(req, stream)
	if err != nil {
		t.Fatalf("ServerStreamingRPC failed: %v", err)
	}

	expectedResponses := []string{
		"Server Streaming RPC response 1",
		"Server Streaming RPC response 2",
		"Server Streaming RPC response 3",
		"Server Streaming RPC response 4",
		"Server Streaming RPC response 5",
	}

	if len(stream.responses) != len(expectedResponses) {
		t.Errorf("Unexpected number of responses. Got: %d, Want: %d", len(stream.responses), len(expectedResponses))
	}

	for i, resp := range stream.responses {
		if resp.Message != expectedResponses[i] {
			t.Errorf("Unexpected response at index %d. Got: %s, Want: %s", i, resp.Message, expectedResponses[i])
		}
	}
}
