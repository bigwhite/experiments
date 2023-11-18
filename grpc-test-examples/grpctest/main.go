package main

import (
	pb "demo/mygrpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC 服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// 注册 MyService 服务
	pb.RegisterMyServiceServer(s, &server{})

	// 启动 gRPC 服务器
	log.Println("Starting gRPC server...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
