package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	pb "rpcdemo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// 重连次数
	retryCount := 5

	// 循环发送请求
	for i := 0; i < retryCount; i++ {
		// 创建 TLS 配置
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true, // 在测试环境中使用，生产环境应设置为 false 并提供证书
		}

		// 创建 gRPC 连接
		conn, err := grpc.Dial(
			"rpc-server.com:443", // 使用 gRPC 服务器的地址和端口
			grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		)
		if err != nil {
			log.Printf("无法连接到服务器: %v", err)
			continue // 继续下一个请求
		}
		defer conn.Close() // 确保连接在每次循环结束时关闭

		// 创建客户端
		client := pb.NewGreetingServiceClient(conn)

		// 构造请求
		request := &pb.GreetingRequest{
			Name: fmt.Sprintf("Client %d", i+1),
		}

		// 发送请求
		response, err := client.Greet(context.Background(), request)
		if err != nil {
			log.Printf("请求失败: %v", err)
			continue // 继续下一个请求
		}

		// 打印响应
		fmt.Printf("响应: %s\n", response.Message)

		// 等待一段时间再发送下一个请求
		time.Sleep(2 * time.Second)
	}
}
