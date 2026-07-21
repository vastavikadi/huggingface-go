package main

import (
	"context"

	"github.com/vastavikadi/huggingface-go"
)

func ExampleEmbeddings(hf_token string) ([]float64, error) {
	client := huggingface.NewClient(huggingface.WithToken(hf_token))

	resp, err := client.Embeddings.Embed(context.Background(), huggingface.EmbedRequest{
		Model: "sentence-transformers/all-MiniLM-L6-v2",
		Input: "The quick brown fox",
	})
	if err != nil {
		return nil, err
	}

	return resp, nil

}
