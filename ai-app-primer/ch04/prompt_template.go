package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template" // 引入 text/template 包
)

// PromptData 定义了填充 Prompt 模板所需的数据结构
// 在实际工程中，这个结构体可能会更复杂，包含更多动态字段
type PromptData struct {
	SystemMessage           string      // 系统/开发者设定的总体指令或角色
	UserQuery               string      // 用户的具体问题
	RetrievedContext        string      // (可选) 通过 RAG 等方式检索到的相关上下文信息
	FewShotExamples         []QAExample // (可选) 用于 Few-Shot Learning 的示例
	OutputFormatInstruction string      // (可选) 对输出格式的明确要求
}

// QAExample 用于定义 Few-Shot Learning 中的问答示例结构
type QAExample struct {
	Question string
	Answer   string
}

// 定义 Prompt 模板字符串。
// 在实际项目中，这个模板通常会存储在一个单独的 .tmpl 文件中，然后被加载进来。
// 注意模板中的 `{{if .FieldName}}...{{end}}` 用于处理可选字段。
// `{{range .ArrayName}}...{{end}}` 用于遍历数组。
const qaPromptTemplate = `{{.SystemMessage}}

{{if .RetrievedContext}}### 相关上下文信息 ###
{{.RetrievedContext}}
{{end}}
{{if .FewShotExamples}}### 问答示例 ###
{{range .FewShotExamples}}
问题：{{.Question}}
回答：{{.Answer}}
---
{{end}}
{{end}}
### 用户提出的问题 ###
{{.UserQuery}}

{{if .OutputFormatInstruction}}### 输出格式要求 ###
{{.OutputFormatInstruction}}
{{end}}
### 你的回答 ###
`

// GenerateQAPrompt 函数接收一个 PromptData 实例，使用模板生成最终的 Prompt 字符串
func GenerateQAPrompt(data PromptData) (string, error) {
	// 1. 创建一个新的模板对象，并给它一个名字 (例如 "qaPrompt")
	tmpl, err := template.New("qaPrompt").Parse(qaPromptTemplate)
	if err != nil {
		// 在实际应用中，模板解析失败通常是一个严重问题，可能需要 panic 或返回更具体的错误
		return "", fmt.Errorf("无法解析 Prompt 模板: %w", err)
	}

	// 2. 创建一个 bytes.Buffer 来接收模板执行后的输出结果
	var filledPrompt bytes.Buffer

	// 3. 执行模板 (Execute)，将 data 对象中的数据填充到模板中
	//    并将结果写入到 filledPrompt 这个 buffer 里
	err = tmpl.Execute(&filledPrompt, data)
	if err != nil {
		return "", fmt.Errorf("无法执行 Prompt 模板并填充数据: %w", err)
	}

	// 4. 从 buffer 中获取最终生成的 Prompt 字符串
	return filledPrompt.String(), nil
}

func main() {
	// 示例用法 1: 一个简单的问答，没有特定上下文或 Few-Shot 示例
	simpleQueryData := PromptData{
		SystemMessage: "你是一个乐于助人的通用问答助手。",
		UserQuery:     "太阳系的中心是什么？",
	}
	prompt1, err := GenerateQAPrompt(simpleQueryData)
	if err != nil {
		log.Fatalf("生成 Prompt 1 失败: %v", err)
	}
	fmt.Println("--- Prompt 1 (简单问答) ---")
	fmt.Println(prompt1)

	// 示例用法 2: 基于提供的上下文回答问题，并要求 JSON 输出
	contextualQueryData := PromptData{
		SystemMessage: "你是一个严谨的知识问答机器人。请仅根据下面提供的“相关上下文信息”来回答用户的问题。如果上下文中没有明确提到，请回答“根据提供的信息无法确定”。",
		UserQuery:     "Go 语言的并发模型是基于线程还是协程？它们之间有何主要区别？",
		RetrievedContext: "Go 语言通过 Goroutines 和 Channels 实现其并发模型。" +
			"Goroutine 是由 Go 运行时管理的轻量级执行单元，与传统的操作系统线程相比，它们的创建和切换成本极低，使得可以轻松创建成千上万个。" +
			"Channels 是 Goroutine 之间进行安全通信和同步的主要机制。",
		OutputFormatInstruction: "请以 JSON 对象的格式回答，包含以下字段：`based_on` (string, 'threads' or 'coroutines'), `key_difference` (string).",
	}
	prompt2, err := GenerateQAPrompt(contextualQueryData)
	if err != nil {
		log.Fatalf("生成 Prompt 2 失败: %v", err)
	}
	fmt.Println("\n--- Prompt 2 (带上下文和 JSON 输出要求) ---")
	fmt.Println(prompt2)

	// 示例用法 3: 使用 Few-Shot 示例来引导翻译风格
	translationData := PromptData{
		SystemMessage: "你是一个专业的英汉翻译引擎。请将用户提供的英文句子翻译成流畅自然的简体中文。",
		UserQuery:     "The quick brown fox jumps over the lazy dog.",
		FewShotExamples: []QAExample{
			{Question: "Hello, world!", Answer: "你好，世界！"},
			{Question: "Knowledge is power.", Answer: "知识就是力量。"},
		},
	}
	prompt3, err := GenerateQAPrompt(translationData)
	if err != nil {
		log.Fatalf("生成 Prompt 3 失败: %v", err)
	}
	fmt.Println("\n--- Prompt 3 (带 Few-Shot 示例的翻译) ---")
	fmt.Println(prompt3)
}
