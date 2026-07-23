package huggingface

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

type ProviderMapping struct {
	ID                       string               `json:"_id"`
	ModelID                  string               `json:"id"`
	InferenceProviderMapping map[string]ModelInfo `json:"inferenceProviderMapping"`
}

type ModelInfo struct {
	Status        string `json:"status"`
	ProviderID    string `json:"providerId"`
	Task          string `json:"task"`
	IsModelAuthor bool   `json:"isModelAuthor"`
}

type FalAIRequest struct {
	Prompt string `json:"prompt"`
}

type Timings struct {
	Inference float64 `json:"inference"`
}

func saveContent(uri string) error {

	u, _ := url.Parse(uri)
	filename := path.Base(u.Path)

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("err from http.Get: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("bad status: ", resp.Status)
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

	return nil
}
