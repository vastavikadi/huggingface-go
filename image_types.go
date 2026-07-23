package huggingface

type ImageGenerationRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`

	Parameters *ImageGenerationParameters `json:"parameters,omitempty"`
}

type ImageGenerationParameters struct {
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty"`
	Width             *int     `json:"width,omitempty"`
	Height            *int     `json:"height,omitempty"`
	Scheduler         *string  `json:"scheduler,omitempty"`
	Seed              *int     `json:"seed,omitempty"`
}

type ImageGenerationResponse struct {
	Image       string
	ContentType string
}

type ImageService struct {
	client *Client
}

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

type FalAIResponse struct {
	Images          []Image `json:"images"`
	Timings         Timings `json:"timings"`
	Seed            int     `json:"seed"`
	HasNSFWConcepts []bool  `json:"has_nsfw_concepts"`
	Prompt          string  `json:"prompt"`
}

type Image struct {
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	ContentType string `json:"content_type"`
}

type Timings struct {
	Inference float64 `json:"inference"`
}
