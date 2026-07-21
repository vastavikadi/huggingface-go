package huggingface

import (
	"net/http"
)

type (
	Client struct {
		token   string
		baseUrl string

		httpClient *http.Client
		Chat       *ChatService
		Embeddings *EmbeddingService
	}
)

func NewClient(opts ...Option) *Client {
	c := &Client{
		httpClient: http.DefaultClient,
		baseUrl:    "https://router.huggingface.co",
	}

	for _, opt := range opts {
		opt(c)
	}

	completions := &ChatCompletionService{
		client: c,
	}

	c.Chat = &ChatService{
		Completions: completions,
	}

	c.Embeddings = &EmbeddingService{
		client: c,
	}

	return c
}
