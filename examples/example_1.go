package main

import (
	"context"
	"fmt"
	"os"

	"github.com/vastavikadi/huggingface-go"
)

const HF_TOKEN = "your_hf_token"

func main() {
	hf_token := os.Getenv("HF_TOKEN")
	if hf_token == "" {
		hf_token = HF_TOKEN
	}
	client := huggingface.NewClient(huggingface.WithToken(hf_token))

	resp, err := client.Chat.Completions.Create(
		context.Background(),
		huggingface.ChatCompletionRequest{
			Model: "openai/gpt-oss-120b",
			Messages: []huggingface.Message{{
				Role:    huggingface.RoleUser,
				Content: "Hello!",
			},
			},
		},
	)
	if err != nil {
		fmt.Println("error chat: ", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
