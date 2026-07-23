package huggingface

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func (s *ImageService) Generate(ctx context.Context, req ImageGenerationRequest) (ImageGenerationResponse, error) {
	if err := req.Validate(); err != nil {
		fmt.Println("err from req.Validate: ", err)
		return ImageGenerationResponse{}, err
	}

	// body := struct {
	// 	Input      string                     `json:"inputs"`
	// 	Parameters *ImageGenerationParameters `json:"parameters,omitempty"`
	// }{
	// 	Input:      req.Prompt,
	// 	Parameters: req.Parameters,
	// }

	var body FalAIRequest

	body.Prompt = req.Prompt

	data, err := s.client.doInferenceRaw(
		ctx,
		req.Model,
		body,
	)
	if err != nil {
		fmt.Println("err from s.client.doInferenceRaw: ", err)
		return ImageGenerationResponse{}, err
	}

	var imageGenResp ImageGenerationResponse

	for _, v := range data.Images {
		imageGenResp.Image = v.URL
		imageGenResp.ContentType = v.ContentType
	}

	err = saveImage(imageGenResp.Image)
	if err != nil {
		fmt.Println("err from saveImage: ", err)
		return ImageGenerationResponse{}, err
	}

	fmt.Printf("imageGenResp %+v", imageGenResp)

	return imageGenResp, nil
}

func (r ImageGenerationRequest) Validate() error {
	if r.Model == "" {
		return errors.New("model is required")
	}

	if r.Prompt == "" {
		return errors.New("prompt is required")
	}

	return nil
}

func saveImage(uri string) error {

	u, _ := url.Parse(uri)
	filename := path.Base(u.Path)

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("err from http.Get: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("err from os.Create: ", err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("err from io.Copy: ", err)
		return err
	}

	fmt.Printf("Downloaded %s", filename)

	return nil
}
