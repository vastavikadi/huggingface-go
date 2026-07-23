package huggingface

import (
	"context"
	"errors"
	"fmt"
)

func (s *VideoService) Generate(ctx context.Context, req VideoGenerationRequest) (VideoGenerationResponse, error) {
	if err := req.Validate(); err != nil {
		fmt.Println("err from req.Validate: ", err)
		return VideoGenerationResponse{}, err
	}

	// body := struct {
	// 	Input      string                     `json:"inputs"`
	// 	Parameters *ImageGenerationParameters `json:"parameters,omitempty"`
	// }{
	// 	Input:      req.Prompt,
	// 	Parameters: req.Parameters,
	// }

	var body FalAIRequest
	var falAiResp FalAIResponseVideo

	body.Prompt = req.Prompt

	err := s.client.doInferenceRaw(
		ctx,
		req.Model,
		body,
		&falAiResp,
	)
	if err != nil {
		fmt.Println("err from s.client.doInferenceRaw: ", err)
		return VideoGenerationResponse{}, err
	}

	var videoGenResp VideoGenerationResponse

	videoGenResp.Video = falAiResp.Videos.URL
	videoGenResp.ContentType = falAiResp.Videos.ContentType

	err = saveContent(videoGenResp.Video)
	if err != nil {
		fmt.Println("err from saveVideo: ", err)
		return VideoGenerationResponse{}, err
	}

	return videoGenResp, nil
}

func (r VideoGenerationRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}

	if r.Prompt == "" {
		return errors.New("prompt is required")
	}

	return nil
}
