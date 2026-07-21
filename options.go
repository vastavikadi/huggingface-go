package huggingface

import "net/http"

type Option func(*Client)

func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseUrl = url
	}
}
