package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// ChatMessage
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	// 省略 ToolCalls 等仅用于工具调用的字段
}

// ChatCompletionRequest
type ChatCompletionRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"` // 包含完整的对话历史
	Temperature float64       `json:"temperature,omitempty"`
	// 使用明确的 max_completion_tokens (如果 API 支持) 或 max_tokens
	MaxTokens int  `json:"max_tokens,omitempty"` // 兼容旧模型或 Ollama
	Stream    bool `json:"stream,omitempty"`
}

// ResponseMessage
type ResponseMessage struct {
	Role    string  `json:"role"`
	Content *string `json:"content"` // 使用指针处理 null
	// 省略 ToolCalls
}

// Choice
type Choice struct {
	Index        int             `json:"index"`
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
	// 省略 logprobs
}

// UsageInfo
type UsageInfo struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ChatCompletionResponse
type ChatCompletionResponse struct {
	ID      string    `json:"id"`
	Object  string    `json:"object"`
	Created int64     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choice  `json:"choices"`
	Usage   UsageInfo `json:"usage"`
	// 省略 system_fingerprint
}

// --- 主函数 ---

const maxTurns = 5 // 控制对话轮数

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	// 对本地 Ollama 等不需要 key 的服务，apiKey 可以为空
	// if apiKey == "" {
	// 	fmt.Println("Warning: OPENAI_API_KEY environment variable not set. Assuming local service.")
	// }

	// --- 配置 API 端点和模型 ---
	// apiURL := "https://api.openai.com/v1/chat/completions" // OpenAI
	// modelID := "gpt-3.5-turbo"
	apiURL := "https://api.deepseek.com/chat/completions"
	modelID := "deepseek-chat" // DeepSeek V3

	// --- 初始化对话历史 ---
	// 开发者消息设定角色和目标
	conversationHistory := []ChatMessage{
		{
			Role:    "system",
			Content: "You are a Go language assistant. Answer questions concisely about Go features.",
		},
	}

	fmt.Printf("Starting a %d-turn conversation with model %s via %s...\n", maxTurns, modelID, apiURL)
	fmt.Println("--------------------------------------------------")

	// --- 创建 HTTP 客户端 ---
	client := &http.Client{Timeout: 60 * time.Second}

	// --- 自动进行多轮对话 ---
	for turn := 1; turn <= maxTurns; turn++ {
		fmt.Printf("--- Turn %d ---\n", turn)

		// 1. 模拟生成用户本轮输入 (基于上一轮的简单逻辑)
		var userPrompt string
		if turn == 1 {
			userPrompt = "What are Go channels used for?"
		} else {
			// 简单追问，实际应用会更复杂
			// 获取上一轮助手的回答来构造问题
			lastAssistantMessage := ""
			if len(conversationHistory) > 0 && conversationHistory[len(conversationHistory)-1].Role == "assistant" {
				lastAssistantMessage = conversationHistory[len(conversationHistory)-1].Content
			}
			if strings.Contains(lastAssistantMessage, "communication") {
				userPrompt = "Can you give a simple code example of channel communication?"
			} else if strings.Contains(lastAssistantMessage, "synchronization") {
				userPrompt = "How does channel synchronization compare to using sync.Mutex?"
			} else {
				// 如果上轮没捕捉到关键词，就问个相关问题
				userPrompt = "What about goroutines? How do they relate to channels?"
			}
		}
		fmt.Printf("User: %s\n", userPrompt)

		// 2. 将用户消息添加到历史记录 (发送前)
		userMessage := ChatMessage{Role: "user", Content: userPrompt}
		conversationHistory = append(conversationHistory, userMessage)

		// 3. 准备 API 请求体 (包含完整历史)
		requestPayload := ChatCompletionRequest{
			Model:       modelID,
			Messages:    conversationHistory, // **关键: 传递了更新后的完整历史**
			Temperature: 0.6,
			MaxTokens:   100, // **限制每轮回复长度**
			Stream:      false,
		}
		requestBodyBytes, err := json.Marshal(requestPayload)
		if err != nil {
			fmt.Printf("[Error] Marshalling request for turn %d: %v\n", turn, err)
			// 出错时移除刚添加的用户消息，避免影响下一轮（如果继续）
			conversationHistory = conversationHistory[:len(conversationHistory)-1]
			continue // 或 break
		}

		// 4. 创建并发送 HTTP 请求
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBodyBytes))
		if err != nil {
			fmt.Printf("[Error] Creating request for turn %d: %v\n", turn, err)
			conversationHistory = conversationHistory[:len(conversationHistory)-1]
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		// 仅在 apiKey 存在时添加 Authorization 头
		if apiKey != "" {
			req.Header.Set("Authorization", "Bearer "+apiKey)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("[Error] Sending request for turn %d: %v\n", turn, err)
			conversationHistory = conversationHistory[:len(conversationHistory)-1]
			continue
		}

		// 5. 处理响应
		func() { // 使用匿名函数方便 defer resp.Body.Close()
			defer resp.Body.Close()
			responseBodyBytes, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				fmt.Printf("[Error] Reading response body for turn %d: %v\n", turn, readErr)
				conversationHistory = conversationHistory[:len(conversationHistory)-1]
				return // 从匿名函数返回
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("[Error] Non-OK status code for turn %d: %d\nResponse: %s\n", turn, resp.StatusCode, string(responseBodyBytes))
				conversationHistory = conversationHistory[:len(conversationHistory)-1]
				return
			}

			var chatResponse ChatCompletionResponse
			unmarshalErr := json.Unmarshal(responseBodyBytes, &chatResponse)
			if unmarshalErr != nil {
				fmt.Printf("[Error] Unmarshalling response for turn %d: %v\n", turn, unmarshalErr)
				fmt.Printf("Raw Response: %s\n", string(responseBodyBytes))
				conversationHistory = conversationHistory[:len(conversationHistory)-1]
				return
			}

			// 6. 提取、打印并存储助手响应
			if len(chatResponse.Choices) > 0 {
				choice := chatResponse.Choices[0]
				assistantContent := ""
				if choice.Message.Content != nil { // 检查 content 是否为 null
					assistantContent = *choice.Message.Content
				}

				fmt.Printf("Assistant: %s\n", assistantContent)
				fmt.Printf("(Finish Reason: %s, Tokens: %d prompt + %d completion = %d total)\n",
					choice.FinishReason,
					chatResponse.Usage.PromptTokens,
					chatResponse.Usage.CompletionTokens,
					chatResponse.Usage.TotalTokens)

				// 将有效的助手响应添加到历史记录 (为下一轮准备)
				if assistantContent != "" || choice.FinishReason == "tool_calls" { // 即使无文本但有工具调用也应记录
					assistantMessage := ChatMessage{Role: "assistant", Content: assistantContent}
					// 如果是工具调用，还需要处理 tool_calls 字段，这里简化处理
					conversationHistory = append(conversationHistory, assistantMessage)
				} else {
					fmt.Printf("[Warning] Assistant response content was empty for turn %d.\n", turn)
					// 如果响应无效，可以选择不添加到历史，或添加一个标记
					// 这里同样移除用户消息，避免空响应影响后续
					conversationHistory = conversationHistory[:len(conversationHistory)-1]
				}
			} else {
				fmt.Printf("[Error] No choices received in response for turn %d.\n", turn)
				conversationHistory = conversationHistory[:len(conversationHistory)-1]
			}
		}() // 立即执行匿名函数

		fmt.Println("--------------------------------------------------")
		// 可以加个短暂休眠，模拟思考时间或避免触发速率限制
		// time.Sleep(1 * time.Second)
	}

	fmt.Println("Conversation finished after 5 turns.")
}
