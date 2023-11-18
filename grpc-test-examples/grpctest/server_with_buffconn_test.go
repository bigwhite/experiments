package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "demo/mygrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func newGRPCServer(t *testing.T) (pb.MyServiceClient, func()) {
	// 创建 bufconn.Listener 作为服务器的监听器
	listener := bufconn.Listen(1024 * 1024)

	// 创建 gRPC 服务器
	srv := grpc.NewServer()

	// 注册服务处理程序
	pb.RegisterMyServiceServer(srv, &server{})

	// 在监听器上启动服务器
	go func() {
		if err := srv.Serve(listener); err != nil {
			t.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 创建 bufconn.Dialer 作为客户端连接
	dialer := func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}

	// 使用 DialContext 和 bufconn.Dialer 创建客户端连接
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}

	// 创建客户端实例
	client := pb.NewMyServiceClient(conn)
	return client, func() {
		err := listener.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		srv.Stop()
	}
}

func TestServerUnaryRPCWithBufConn(t *testing.T) {
	client, shutdown := newGRPCServer(t)
	defer shutdown()

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := client.UnaryRPC(context.Background(), tt.requestMessage)
			if err != nil {
				t.Fatalf("UnaryRPC failed: %v", err)
			}

			if resp.Message != tt.expectedResp.Message {
				t.Errorf("Unexpected response. Got: %s, Want: %s", resp.Message, tt.expectedResp.Message)
			}
		})
	}
}

func TestServerServerStreamingRPCWithBufConn(t *testing.T) {
	client, shutdown := newGRPCServer(t)
	defer shutdown()

	req := &pb.RequestMessage{
		Message: "Test message",
	}

	stream, err := client.ServerStreamingRPC(context.Background(), req)
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

	gotResponses := []string{}

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		gotResponses = append(gotResponses, resp.Message)
	}

	if len(gotResponses) != len(expectedResponses) {
		t.Errorf("Unexpected number of responses. Got: %d, Want: %d", len(gotResponses), len(expectedResponses))
	}

	for i, resp := range gotResponses {
		if resp != expectedResponses[i] {
			t.Errorf("Unexpected response at index %d. Got: %s, Want: %s", i, resp, expectedResponses[i])
		}
	}
}
