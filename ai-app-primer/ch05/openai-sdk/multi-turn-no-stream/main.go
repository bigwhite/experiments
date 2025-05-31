package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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

	// 2. 客户端初始化 (同上)
	client := openai.NewClient(option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL))
	ctx := context.Background()

	// 3. 初始化对话历史
	// 创建一个 openai.ChatCompletionMessageParamUnion 类型的切片来存储历史。
	// 通常以一个 openai.SystemMessage 开始，设定助手的角色或行为。
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage("你是一个乐于助人的 Go 语言编程助手。"),
	}

	fmt.Println("开始与 Go 助手对话 (输入 'quit' 退出):")
	// 使用 bufio.NewReader 读取用户输入。
	reader := bufio.NewReader(os.Stdin)

	// 4. 进入对话循环
	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// 输入 "quit" 退出循环。
		if strings.ToLower(userInput) == "quit" {
			fmt.Println("再见!")
			break
		}

		// 5. 将用户输入添加到历史记录
		// 使用 openai.UserMessage 将用户输入包装并追加到 messages 切片。
		messages = append(messages, openai.UserMessage(userInput))

		// 6. 构建包含完整历史的请求
		params := openai.ChatCompletionNewParams{
			Model: "deepseek-chat",
			// 关键：将包含所有历史的 messages 切片传递给 Messages 字段。
			Messages: messages,
		}

		fmt.Println("助手正在思考...")

		// 7. 发起 API 请求 (非流式)
		completion, err := client.Chat.Completions.New(ctx, params)
		// 8. 处理 API 错误
		if err != nil {
			var apiErr *openai.Error
			if errors.As(err, &apiErr) {
				fmt.Printf("API 错误: Status=%d Type=%s Message=%s\n", apiErr.StatusCode, apiErr.Type, apiErr.Message)
			} else {
				fmt.Printf("API 错误: %v\n", err)
			}
			// 可选：如果调用失败，从历史记录中移除刚才添加的用户消息，避免错误累积。
			messages = messages[:len(messages)-1]
			continue // 继续下一次循环，等待用户再次输入。
		}

		// 9. 处理并添加助手响应到历史记录
		if len(completion.Choices) > 0 {
			// 获取助手的响应消息结构体 (openai.ChatCompletionMessage)。
			assistantResponseMsg := completion.Choices[0].Message
			fmt.Printf("Assistant: %s\n", assistantResponseMsg.Content)

			// 关键步骤：将助手的响应消息转换回参数类型，并添加到历史记录中。
			// 使用 assistantResponseMsg.ToParam() 方法！
			messages = append(messages, assistantResponseMsg.ToParam())

		} else {
			fmt.Println("Assistant: 我暂时没有回应。")
			// 可选：如果助手没有回应，也移除最后的用户消息。
			messages = messages[:len(messages)-1]
		}

		// 10. 可选：历史记录截断逻辑
		// 在实际应用中，需要检查 messages 的长度（或累计 token 数）
		// 并根据需要移除旧的消息，以防止超出模型的上下文窗口限制。
		// const maxHistoryItems = 10 // 例如保留最近 10 条（含系统消息）
		// if len(messages) > maxHistoryItems {
		//    // 保留第一条（系统消息）和最后 maxHistoryItems-1 条
		//    messages = append(messages[:1], messages[len(messages)-(maxHistoryItems-1):]...)
		// }
	}
}
