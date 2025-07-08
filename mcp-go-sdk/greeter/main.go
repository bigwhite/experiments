package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HiParams 定义了工具的输入参数，强类型保证
type HiParams struct {
	Name string `json:"name"`
}

// SayHi 是工具的具体实现
func SayHi(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[HiParams]) (*mcp.CallToolResultFor[any], error) {
	resultText := fmt.Sprintf("Hi %s, welcome to the Go MCP world!", params.Arguments.Name)
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: resultText},
		},
	}, nil
}

func main() {
	// 1. 创建 Server 实例
	server := mcp.NewServer("greeter-server", "1.0.0", nil)

	// 2. 添加工具
	// NewServerTool 利用泛型和反射自动生成输入 schema
	server.AddTools(
		mcp.NewServerTool("greet", "Say hi to someone", SayHi),
	)

	// 3. 通过 StdioTransport 运行服务，它会监听标准输入/输出
	log.Println("Greeter server running over stdio...")
	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatalf("Server run failed: %v", err)
	}
}
