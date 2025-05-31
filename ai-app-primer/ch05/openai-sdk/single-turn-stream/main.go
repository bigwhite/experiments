package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	// 1. 环境准备
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("错误：未设置 OPENAI_API_KEY 环境变量。")
	}
	baseURL := os.Getenv("OPENAI_API_BASE")
	if baseURL == "" {
		log.Fatal("Error: OPENAI_API_BASE environment variable not set.")
	}

	// 2. 客户端初始化
	client := openai.NewClient(option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL))
	ctx := context.Background()

	// 3. 构建流式请求参数
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("请简单介绍一下 Go 语言的主要特性。"),
		},
		Model:       "deepseek-chat",
		MaxTokens:   openai.Int(200),
		Temperature: openai.Float(0.7),
	}

	fmt.Println("正在向 OpenAI 发送流式请求...")
	fmt.Println("--- OpenAI 流式响应 ---")

	// 4. 调用 Chat Completion API (流式)
	// 使用 client.Chat.Completions.NewStreaming 方法发起流式请求。
	stream := client.Chat.Completions.NewStreaming(ctx, params)

	// 5. 循环接收和处理流数据块
	// 使用 for stream.Next() 循环，直到流结束或发生错误。
	for stream.Next() {
		// 获取当前数据块 chunk (类型为 openai.ChatCompletionChunk)。
		chunk := stream.Current()
		// 检查当前块的 Choices 是否有内容。
		if len(chunk.Choices) > 0 {
			// 从 chunk.Choices[0].Delta.Content 获取增量的文本片段并打印。
			// Delta 表示这是相对于上一个块的变化。
			fmt.Printf(chunk.Choices[0].Delta.Content)
		}
	}

	// 6. 检查流处理过程中的错误
	// 循环结束后，必须调用 stream.Err() 检查是否有错误发生。
	if err := stream.Err(); err != nil {
		log.Printf("\n流处理错误: %v\n", err)
		// 注意：新 SDK 不再使用 io.EOF 判断流结束，循环自然结束即表示流完成。
		return
	}

	fmt.Println("\n流处理完成。")
	// 注意：流式响应通常不直接在每个块中提供总的 Usage 信息。
	// 可能需要使用 Accumulator 或在流结束后通过其他方式获取。
}
