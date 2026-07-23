package main

import (
	"context"

	"github.com/vastavikadi/huggingface-go"
)

func ExampleVideoGen(hf_token string) (string, error) {
	client := huggingface.NewClient(huggingface.WithToken(hf_token))

	resp, err := client.VideoGen.Generate(context.Background(), huggingface.VideoGenerationRequest{
		Model:  "Wan-AI/Wan2.2-TI2V-5B",
		Prompt: "Astronaut riding a horse",
	})

	if err != nil {
		return "", err
	}

	return resp.Video, nil
}
