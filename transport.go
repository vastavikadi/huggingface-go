package huggingface

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
