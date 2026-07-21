package main

import (
	"context"
	"fmt"

	"github.com/vastavikadi/huggingface-go"
)

func ExampleChat(hf_token string) (string, error) {
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
		return "", fmt.Errorf("error chat: %v", err)
	}

	return (resp.Choices[0].Message.Content), nil
}
