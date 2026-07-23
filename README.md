<div align="center">

# HuggingFace-Go

### An idiomatic Go SDK for the Hugging Face Inference API

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/vastavikadi/huggingface-go?style=flat-square)](https://github.com/vastavikadi/huggingface-go)

A Go SDK for the Hugging Face Inference API.

HuggingFace-Go provides an idiomatic Go interface for interacting with Hugging Face models without dealing with raw HTTP requests.

*Inspired by Python's `huggingface_hub.InferenceClient`*

</div>

## Features

- Chat Completions
- Text Embeddings
- Idiomatic Go API
- Context-aware requests
- Configurable HTTP client
- Functional options
- Lightweight with minimal dependencies
- Supports image and video generation via FalAI provider

> NOTE: huggingface has deprecated the `hf-inference` endpoint for image and video generation. This SDK uses the `fal-ai` provider for these tasks. There are many other providers available for image and video generation, but this SDK currently only supports `fal-ai`. If you want to use other providers, please open an issue or submit a PR.

## Installation

```bash
go get github.com/vastavikadi/huggingface-go
```

## Quick Start

```go
client := huggingface.NewClient(
    huggingface.WithToken(os.Getenv("HF_TOKEN")),
)

resp, err := client.Chat.Completions.Create(
    context.Background(),
    huggingface.CreateChatCompletionRequest{
        Model: "openai/gpt-oss-120b",
        Messages: []huggingface.Message{
            {
                Role:    huggingface.RoleUser,
                Content: "Hello!",
            },
        },
    },
)
```

### Embeddings

```go
embedding, err := client.Embed(
    context.Background(),
    huggingface.EmbedRequest{
        Model: "sentence-transformers/all-MiniLM-L6-v2",
        Input: "The quick brown fox",
    },
)
```

### Image Generation

```go
imageResp, err := client.Image.Generate(
    context.Background(),
    huggingface.ImageGenerationRequest{
        Model:  "black-forest-labs/FLUX.1-dev",
        Prompt: "A futuristic cityscape at sunset",
    },
)
```

### Video Generation

```go
videoResp, err := client.Video.Generate(
    context.Background(),
    huggingface.VideoGenerationRequest{
        Model:  "Wan-AI/Wan2.2-TI2V-5B",
        Prompt: "A futuristic cityscape at sunset",
    },
)
```

## Roadmap

- [x] Chat Completions
- [x] Text Embeddings
- [x] Image Generation
- [x] Video Generation
- [ ] Streaming
- [ ] Audio
- [ ] Vision
- [ ] Additional Hugging Face inference tasks

## Examples

Explore the `examples/` directory for comprehensive usage examples:

| File | Purpose |
|------|---------|
| `example_chat.go` | 💬 Chat completions example |
| `example_embeddings.go` | 🧮 Text embedding generation example |
| `example_imageGen.go` | 🖼️ Image generation example |
| `example_videoGen.go` | 🎬 Video generation example |
| `main.go` | 🎯 Execute the examples |


### Running Examples

Execute any example with a single command:

```bash
go run ./examples
```

## Contributing
- We'd love your contributions! Whether it's bug fixes, new features, or documentation improvements, here's how to get started:

### Steps to Contribute
- Fork the repository

```bash
git clone https://github.com/your-username/huggingface-go.git
```

- Create a feature branch

```bash
git checkout -b feature/amazing-feature
```

- Make your changes and commit

```bash
git commit -m 'Add amazing feature'
```

- Push to your branch

```bash
git push origin feature/amazing-feature
```

- Open a Pull Request on GitHub

## Guidelines
- Follow Go code standards
- Add tests for new features
- Update documentation as needed
- Keep commits clean and descriptive

## License

MIT