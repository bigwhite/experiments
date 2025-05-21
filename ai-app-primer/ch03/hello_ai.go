package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 请求结构体 (简化版，只包含核心字段)
type ChatCompletionRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 响应结构体 (简化版，只关注核心信息)
type ChatCompletionResponse struct {
	ID      string    `json:"id"`
	Choices []Choice  `json:"choices"`
	Usage   UsageInfo `json:"usage"`
	// 为简化，省略了 object, created, model, system_fingerprint 等
}

type Choice struct {
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
	// 为简化，省略了 index, logprobs 等
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"` // 注意：在实际应用中，如果 content 可能为 null (如 tool_calls 时)，应使用 *string
}

type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable not set.")
		return
	}

	// DeepSeek API 配置
	apiURL := "https://api.deepseek.com/chat/completions" // 或者 "https://api.deepseek.com/v1/chat/completions"
	modelID := "deepseek-chat"                            // 使用 DeepSeek 的一个聊天模型

	// 构造请求体
	requestPayload := ChatCompletionRequest{
		Model: modelID,
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a friendly and helpful assistant."}, // System (or Developer) message
			{Role: "user", Content: "Hello, AI! How are you today?"},               // User message
		},
	}

	requestBodyBytes, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Printf("Error marshalling request: %v\n", err)
		return
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Non-OK status code: %d\nResponse: %s\n", resp.StatusCode, string(responseBodyBytes))
		return
	}

	// 解析 JSON 响应
	var chatResponse ChatCompletionResponse
	err = json.Unmarshal(responseBodyBytes, &chatResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		fmt.Printf("Raw Response: %s\n", string(responseBodyBytes))
		return
	}

	// 打印关键信息
	fmt.Println("--- AI Response ---")
	if len(chatResponse.Choices) > 0 {
		fmt.Printf("ID: %s\n", chatResponse.ID)
		fmt.Printf("Assistant: %s\n", chatResponse.Choices[0].Message.Content)
		fmt.Printf("Finish Reason: %s\n", chatResponse.Choices[0].FinishReason)
		fmt.Printf("Usage: PromptTokens=%d, CompletionTokens=%d, TotalTokens=%d\n",
			chatResponse.Usage.PromptTokens,
			chatResponse.Usage.CompletionTokens,
			chatResponse.Usage.TotalTokens,
		)
	} else {
		fmt.Println("No choices received.")
	}
}
