package main

import (
	"context"
	"log"
	"net"

	proto "github.com/bigwhite/grpc/proto"
	"google.golang.org/grpc"
	//pb "github.com/gogo/protobuf/proto"
)

type PublishServiceImpl struct{}

func (p *PublishServiceImpl) Publish(ctx context.Context, m *proto.Message) (*proto.Response, error) {
	//log.Println("recv a rpc request:", *m)
	ack := &proto.Response{
		Code: 0,
		Msg:  "ok",
	}
	return ack, nil
}

// https://pkg.go.dev/google.golang.org/grpc
func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterPublishServiceServer(grpcServer, new(PublishServiceImpl))

	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	grpcServer.Serve(lis)
}
