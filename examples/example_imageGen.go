package main

import (
	"context"

	"github.com/vastavikadi/huggingface-go"
)

func ExampleImageGen(hf_token string) (string, error) {
	client := huggingface.NewClient(huggingface.WithToken(hf_token))

	resp, err := client.ImageGen.Generate(context.Background(), huggingface.ImageGenerationRequest{
		Model:  "black-forest-labs/FLUX.1-dev",
		Prompt: "Astronaut riding a horse",
	})
	if err != nil {
		return "", nil
	}

	return resp.Image, nil
}
