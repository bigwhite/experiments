package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// serverConfig 结构体用于管理不同 MCP 服务的连接信息
type serverConfig struct {
	ServerCmd string // 用于 stdio 服务
	HTTPAddr  string // 用于 http 服务
}

// toolRegistry 映射对 LLM 友好的工具别名到其服务配置
var toolRegistry = map[string]serverConfig{
	"greet":      {ServerCmd: "go run ../greeter/main.go"},
	"add":        {HTTPAddr: "http://localhost:8080/math"},
	"list_files": {ServerCmd: "go run ../fileserver/main.go"},
	"read_file":  {ServerCmd: "go run ../fileserver/main.go"},
}

// invokeMCPTool 是 Agent 的核心函数，负责直接与 MCP 服务通信
func invokeMCPTool(toolAlias string, arguments map[string]interface{}) (string, error) {
	config, ok := toolRegistry[toolAlias]
	if !ok {
		return "", fmt.Errorf("unknown tool alias: %s", toolAlias)
	}

	// 1. 将 LLM 友好的别名和参数，转换为真正的 MCP 请求
	mcpToolName := toolAlias
	mcpArguments := arguments
	if toolAlias == "list_files" {
		mcpToolName = "resources/read"
		mcpArguments = map[string]interface{}{"uri": "mcp://fs/list"}
	} else if toolAlias == "read_file" {
		mcpToolName = "resources/read"
		if filename, ok := arguments["filename"].(string); ok {
			mcpArguments = map[string]interface{}{"uri": "file:///" + filename}
		} else {
			return "", fmt.Errorf("tool 'read_file' requires a 'filename' argument")
		}
	}

	// 2. 创建 MCP 客户端实例
	client := mcp.NewClient("go-agent", "1.0", nil)

	// 3. 根据配置选择并创建 Transport
	var transport mcp.Transport
	if config.ServerCmd != "" {
		cmdParts := strings.Fields(config.ServerCmd)
		transport = mcp.NewCommandTransport(exec.Command(cmdParts[0], cmdParts[1:]...))
	} else {
		transport = mcp.NewStreamableClientTransport(config.HTTPAddr, nil)
	}

	// 4. 授权客户端访问本地文件系统（仅对文件服务调用有效）
	client.AddRoots(&mcp.Root{URI: "file://./"})

	// 5. 连接到服务器，建立会话
	ctx := context.Background()
	session, err := client.Connect(ctx, transport)
	if err != nil {
		return "", fmt.Errorf("failed to connect to MCP server for tool %s: %w", toolAlias, err)
	}
	defer session.Close() // 每次调用都是一个独立的会话，确保关闭

	// 6. 执行调用并处理结果
	var resultText string
	if mcpToolName == "resources/read" {
		res, err := session.ReadResource(ctx, &mcp.ReadResourceParams{
			URI: mcpArguments["uri"].(string),
		})
		if err != nil {
			return "", fmt.Errorf("ReadResource failed: %w", err)
		}
		var sb strings.Builder
		for _, c := range res.Contents {
			sb.WriteString(c.Text)
		}
		resultText = sb.String()
	} else {
		res, err := session.CallTool(ctx, &mcp.CallToolParams{
			Name:      mcpToolName,
			Arguments: mcpArguments,
		})
		if err != nil {
			return "", fmt.Errorf("CallTool failed: %w", err)
		}
		if res.IsError {
			return "", fmt.Errorf("tool execution failed: %s", res.Content[0].(*mcp.TextContent).Text)
		}
		resultText = res.Content[0].(*mcp.TextContent).Text
	}

	return resultText, nil
}

func main() {
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		log.Fatal("DEEPSEEK_API_KEY environment variable not set.")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL("https://api.deepseek.com/v1"),
	)

	// 为所有工具使用合法的名称，特别是为 `resources/read` 创建别名
	tools := []openai.ChatCompletionToolParam{
		{
			Function: openai.FunctionDefinitionParam{
				Name:        "greet",
				Description: openai.String("Say hi to someone."),
				Parameters: openai.FunctionParameters{
					"type": "object", "properties": map[string]interface{}{"name": map[string]string{"type": "string", "description": "Name of the person to greet"}}, "required": []string{"name"},
				},
			},
		},
		{
			Function: openai.FunctionDefinitionParam{
				Name:        "add",
				Description: openai.String("Add two integers."),
				Parameters: openai.FunctionParameters{
					"type": "object", "properties": map[string]interface{}{"A": map[string]string{"type": "integer"}, "B": map[string]string{"type": "integer"}}, "required": []string{"A", "B"},
				},
			},
		},
		{
			Function: openai.FunctionDefinitionParam{
				Name:        "list_files",
				Description: openai.String("List all non-directory files in the current project directory."),
				Parameters:  openai.FunctionParameters{"type": "object", "properties": map[string]interface{}{}},
			},
		},
		{
			Function: openai.FunctionDefinitionParam{
				Name:        "read_file",
				Description: openai.String("Read the content of a specific file."),
				Parameters: openai.FunctionParameters{
					"type": "object", "properties": map[string]interface{}{"filename": map[string]string{"type": "string", "description": "The name of the file to read."}}, "required": []string{"filename"},
				},
			},
		},
	}

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage("You are a helpful assistant with access to local tools. You must call tools by using the tool_calls response format. Don't make assumptions about what values to plug into functions. Ask for clarification if a user request is ambiguous."),
		openai.UserMessage("Hi, can you greet my friend Alex, add 5 and 7, and then list the files in my project?"),
	}

	ctx := context.Background()

	for i := 0; i < 5; i++ {
		log.Println("--- Sending request to DeepSeek ---")

		resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{Model: "deepseek-chat", Messages: messages, Tools: tools})
		if err != nil {
			log.Fatalf("ChatCompletion error: %v\n", err)
		}
		if len(resp.Choices) == 0 {
			log.Fatal("No choices returned from API")
		}

		msg := resp.Choices[0].Message
		messages = append(messages, msg.ToParam())

		if msg.ToolCalls != nil {
			for _, toolCall := range msg.ToolCalls {
				functionName := toolCall.Function.Name
				var arguments map[string]interface{}
				if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &arguments); err != nil {
					log.Fatalf("Failed to unmarshal function arguments: %v", err)
				}

				log.Printf("--- LLM wants to call tool: %s with args: %v ---\n", functionName, arguments)

				// 直接调用我们的 Go 函数，该函数内建了 MCP 客户端逻辑
				toolResult, err := invokeMCPTool(functionName, arguments)
				if err != nil {
					log.Printf("Tool call failed: %v\n", err)
					toolResult = fmt.Sprintf("Error executing tool: %v", err)
				}

				log.Printf("--- Tool result: ---\n%s\n---------------------\n", toolResult)

				messages = append(messages, openai.ToolMessage(toolResult, toolCall.ID))
			}
			continue
		}

		log.Println("--- Final response from LLM ---")
		log.Println(msg.Content)
		return
	}
	log.Println("Reached max conversation turns.")
}
