package huggingface

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) do(
	ctx context.Context,
	method string,
	path string,
	reqBody any,
	respBody any,
) error {

	var body io.Reader

	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseUrl+path, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("huggingface: status %d: %s", res.StatusCode, string(data))
	}

	if respBody != nil {
		if err := json.Unmarshal(data, respBody); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) doInference(
	ctx context.Context,
	model string,
	task string,
	req any,
	resp any,
) error {
	path := fmt.Sprintf("/hf-inference/models/%s/pipeline/%s", url.PathEscape(model), task)

	return c.do(
		ctx, http.MethodPost, path, req, resp,
	)
}

func (c *Client) doInferenceRaw(ctx context.Context, model string, reqBody any) (FalAIResponse, error) {

	inferenceRes, err := GetProviders(model)
	if err != nil {
		fmt.Println("err from inferenceRes: ", err)
		return FalAIResponse{}, err
	}
	var provider string
	var providerId string

	for i, v := range inferenceRes.InferenceProviderMapping {
		provider = i
		providerId = v.ProviderID
		break
	}

	path := fmt.Sprintf("/%s/%s", provider, providerId)

	var body io.Reader

	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			fmt.Println("err from json.Marshal(reqBody): ", err)
			return FalAIResponse{}, err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl+path, body)
	if err != nil {
		fmt.Println("err from http.NewRequestWithContext: ", err)
		return FalAIResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("err from c.httpClient.Do: ", err)
		return FalAIResponse{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err from io.ReadAll: ", err)
		return FalAIResponse{}, err
	}

	if res.StatusCode >= 400 {
		fmt.Println("err from inference response: ", fmt.Errorf("huggingface: status %d: %s", res.StatusCode, string(data)))
		return FalAIResponse{}, fmt.Errorf("huggingface: status %d: %s", res.StatusCode, string(data))
	}

	var falAiResp FalAIResponse

	err = json.Unmarshal(data, &falAiResp)
	if err != nil {
		fmt.Println("err from json.Unmarshal: ", err)
		return FalAIResponse{}, err
	}

	return falAiResp, nil
}

// providers are required for image generations
// most of the hf-inference models are deprecated
func GetProviders(model string) (ProviderMapping, error) {
	var inferenceResp ProviderMapping

	pathForInferenceProvider := fmt.Sprintf("https://huggingface.co/api/models/%s?expand[]=inferenceProviderMapping", model)

	req, err := http.NewRequest(http.MethodGet, pathForInferenceProvider, nil)
	if err != nil {
		fmt.Println("err from http.NewRequest: ", err)
		return ProviderMapping{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err from http.DefaultClient.Do: ", err)
		return ProviderMapping{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err from io.ReadAll: ", err)
		return ProviderMapping{}, err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("err from providers response: ", fmt.Errorf("status %d: %s", res.StatusCode, string(data)))
		return ProviderMapping{}, fmt.Errorf("status %d: %s", res.StatusCode, string(data))
	}

	if err = json.Unmarshal(data, &inferenceResp); err != nil {
		fmt.Println("err from json.Unmarshal: ", err)
		return ProviderMapping{}, err
	}
	return inferenceResp, nil
}
