package huggingface

type ChatService struct {
	Completions *ChatCompletionService
}

type ChatCompletionService struct {
	client *Client
}


