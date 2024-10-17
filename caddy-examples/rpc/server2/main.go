package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "rpcdemo/proto"

	"google.golang.org/grpc"
)

// 服务器结构体
type server struct {
	pb.UnimplementedGreetingServiceServer
}

// 实现 Greet 方法
func (s *server) Greet(ctx context.Context, req *pb.GreetingRequest) (*pb.GreetingReply, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	fmt.Println("Recv req from:", req.Name)
	return &pb.GreetingReply{Message: message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9008")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetingServiceServer(grpcServer, &server{})

	fmt.Println("gRPC server is running on port 9008...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
