package main

import (
	"context"
	"testing"

	pb "demo/mygrpc"
)

func TestServerUnaryRPC(t *testing.T) {
	s := &server{}

	req := &pb.RequestMessage{
		Message: "Test message",
	}

	resp, err := s.UnaryRPC(context.Background(), req)
	if err != nil {
		t.Fatalf("UnaryRPC failed: %v", err)
	}

	expectedResp := &pb.ResponseMessage{
		Message: "Unary RPC response",
	}

	if resp.Message != expectedResp.Message {
		t.Errorf("Unexpected response. Got: %s, Want: %s", resp.Message, expectedResp.Message)
	}
}

func TestServerUnaryRPCs(t *testing.T) {
	tests := []struct {
		name           string
		requestMessage *pb.RequestMessage
		expectedResp   *pb.ResponseMessage
	}{
		{
			name: "Test Case 1",
			requestMessage: &pb.RequestMessage{
				Message: "Test message",
			},
			expectedResp: &pb.ResponseMessage{
				Message: "Unary RPC response",
			},
		},
		// Add more test cases as needed
	}

	s := &server{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.UnaryRPC(context.Background(), tt.requestMessage)
			if err != nil {
				t.Fatalf("UnaryRPC failed: %v", err)
			}

			if resp.Message != tt.expectedResp.Message {
				t.Errorf("Unexpected response. Got: %s, Want: %s", resp.Message, tt.expectedResp.Message)
			}
		})
	}
}
