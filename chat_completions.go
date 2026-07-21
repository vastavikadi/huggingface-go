package huggingface

import (
	"context"
	"net/http"
)

func (s *ChatCompletionService) Create(ctx context.Context, req ChatCompletionRequest) (*ChatCompletionResponse, error) {
	var resp ChatCompletionResponse

	err := s.client.do(
		ctx, http.MethodPost, "/v1/chat/completions", req, &resp,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
