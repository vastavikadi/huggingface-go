package huggingface

type EmbedRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}
type EmbedBatchRequest struct {
	Model  string   `json:"model"`
	Inputs []string `json:"inputs"`
}

type Embedding struct {
	Object    string    `json:"object"`
	Index     int       `json:"index"`
	Embedding []float32 `json:"embedding"`
}

type EmbeddingService struct {
	client *Client
}
