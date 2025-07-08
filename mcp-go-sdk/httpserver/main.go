package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// HiParams 和 SayHi 函数与场景一相同
type HiParams struct {
	Name string `json:"name"`
}

func SayHi(ctx context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[HiParams]) (*mcp.CallToolResultFor[any], error) {
	resultText := fmt.Sprintf("Hi %s, this response is from the HTTP server!", params.Arguments.Name)
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: resultText}},
	}, nil
}

// AddParams 和 Add 工具的实现
type AddParams struct{ A, B int }

func Add(_ context.Context, _ *mcp.ServerSession, params *mcp.CallToolParamsFor[AddParams]) (*mcp.CallToolResultFor[any], error) {
	result := params.Arguments.A + params.Arguments.B
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("The sum is: %d", result)}},
	}, nil
}

func main() {
	// 1. 创建 Greeter 服务实例
	greeterServer := mcp.NewServer("greeter-service", "1.0", nil)
	greeterServer.AddTools(mcp.NewServerTool("greet", "Say hi", SayHi))

	// 2. 创建 Math 服务实例
	mathServer := mcp.NewServer("math-service", "1.0", nil)
	mathServer.AddTools(mcp.NewServerTool("add", "Add two integers", Add))

	// 3. 创建 StreamableHTTPHandler
	handler := mcp.NewStreamableHTTPHandler(func(request *http.Request) *mcp.Server {
		log.Printf("Routing request for URL: %s\n", request.URL.Path)
		switch request.URL.Path {
		case "/greeter":
			return greeterServer
		case "/math":
			return mathServer
		default:
			return nil // 返回 nil 将导致 404 Not Found
		}
	}, nil)

	// 4. 启动标准的 Go HTTP 服务器
	addr := ":8080"
	log.Printf("Multi-service MCP server listening at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
