package huggingface

type VideoGenerationRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`

	Parameters *VideoGenerationParameters `json:"parameters,omitempty"`
}

type VideoGenerationParameters struct {
	NegativePrompt    *string  `json:"negative_prompt,omitempty"`
	GuidanceScale     *float64 `json:"guidance_scale,omitempty"`
	NumInferenceSteps *int     `json:"num_inference_steps,omitempty"`
	Width             *int     `json:"width,omitempty"`
	Height            *int     `json:"height,omitempty"`
	Scheduler         *string  `json:"scheduler,omitempty"`
	Seed              *int     `json:"seed,omitempty"`
}

type VideoGenerationResponse struct {
	Video       string
	ContentType string
}

type VideoService struct {
	client *Client
}

type FalAIResponseVideo struct {
	Videos Video `json:"video"`
	Seed   int     `json:"seed"`
	Prompt string  `json:"prompt"`
}

type Video struct {
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
	FileSize    int    `json:"file_size"`
}
