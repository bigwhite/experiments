package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable not set.")
	}

	baseURL := os.Getenv("OPENAI_API_BASE")
	if baseURL == "" {
		log.Fatal("Error: OPENAI_API_BASE environment variable not set.")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey),
		option.WithEndpoint(baseURL))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("deepseek-chat")
	cs := model.StartChat()
	res, err := cs.SendMessage(ctx, genai.Text("你好，请用Go语言写一个简单的hello world程序"))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(res)

}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
