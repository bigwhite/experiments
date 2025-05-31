package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	// 1. 从环境变量获取 API 密钥和Base URL
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable not set.")
	}
	baseURL := os.Getenv("OPENAI_API_BASE")
	if baseURL == "" {
		log.Fatal("Error: OPENAI_API_BASE environment variable not set.")
	}

	// 2. 创建 OpenAI 客户端
	client := openai.NewClient(option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL))
	ctx := context.Background()

	// 3. 构建请求参数，使用 openai.ChatCompletionNewParams 结构体定义请求。
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			// Use message constructor helpers
			openai.UserMessage("你好，请用Go语言写一个简单的hello world程序。"),
		},
		Model: "deepseek-chat",
		// Optional parameters using helpers
		MaxTokens:   openai.Int(100),
		Temperature: openai.Float(0.7),
	}

	fmt.Println("Sending request to OpenAI...")

	// 4. 调用 Chat Completion API (非流式)
	// 使用 client.Chat.Completions.New 方法发起请求。
	completion, err := client.Chat.Completions.New(ctx, params)
	// 5. 错误处理
	if err != nil {
		// 使用 errors.As 检查是否为 OpenAI API 特有的错误 *openai.Error。
		var apiErr *openai.Error
		if errors.As(err, &apiErr) {
			log.Printf("OpenAI API error: Status=%d Type=%s Code=%v Param=%s Message=%s\n",
				apiErr.StatusCode, apiErr.Type, apiErr.Code, apiErr.Param, apiErr.Message)
		} else {
			log.Printf("Error calling OpenAI API: %v\n", err)
		}
		return // Exit if error occurs
	}

	// 6. 处理成功的响应
	// 检查 completion.Choices 是否包含内容。
	if len(completion.Choices) > 0 {
		// Access the response content
		fmt.Println("--- OpenAI Response ---")
		fmt.Println(completion.Choices[0].Message.Content)
		fmt.Println("-----------------------")
		// Access usage information (structure might be slightly different, check docs if needed)
		// Assuming Usage struct still exists and has these fields
		fmt.Printf("Usage - Prompt Tokens: %d, Completion Tokens: %d, Total Tokens: %d\n",
			completion.Usage.PromptTokens, completion.Usage.CompletionTokens, completion.Usage.TotalTokens)
	} else {
		fmt.Println("No response choices received.")
	}
}
