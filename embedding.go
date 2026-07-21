package huggingface

import (
	"context"
	"errors"
)

func (s *EmbeddingService) Embed(ctx context.Context, req EmbedRequest) ([]float64, error) {
	var embedding []float64

	if err := req.Validate(); err != nil {
		return nil, err
	}

	body := struct {
		Input string `json:"inputs"`
	}{
		Input: req.Input,
	}

	err := s.client.doInference(
		ctx,
		req.Model,
		"feature-extraction",
		body,
		&embedding,
	)

	if err != nil {
		return nil, err
	}

	return embedding, nil
}

func (r EmbedRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}

	if r.Input == "" {
		return errors.New("inputs are required")
	}

	return nil
}

func (s *EmbeddingService) EmbedBatch(ctx context.Context, req EmbedBatchRequest) ([][]float64, error) {
	var embedding [][]float64

	body := struct {
		Inputs []string `json:"inputs"`
	}{
		Inputs: req.Inputs,
	}

	err := s.client.doInference(
		ctx,
		req.Model,
		"feature-extraction",
		body,
		&embedding,
	)

	if err != nil {
		return nil, err
	}

	return embedding, nil
}
