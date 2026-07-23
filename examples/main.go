package main

import (
	"fmt"
	"os"
)

// const HF_TOKEN = ""

func main() {
	hf_token := os.Getenv("HF_TOKEN")
	// if hf_token == "" {
	// 	hf_token = HF_TOKEN
	// }

	// chat api demo
	res, err := ExampleChat(hf_token)
	if err != nil {
		fmt.Println("error generating chat: ", err)
	}

	fmt.Println("res from chat: ", res)

	// Embeddings api demo
	resp, err := ExampleEmbeddings(hf_token)
	if err != nil {
		fmt.Println("error creating embeddings: ", err)
	}

	fmt.Println("resp creating embeddings: ", resp)

	// ImageGen api demo
	imgPath, err := ExampleImageGen(hf_token)
	if err != nil {
		fmt.Println("error creating image: ", err)
	}

	fmt.Println("resp creating images: ", imgPath)

	// VideoGen api demo
	videoPath, err := ExampleVideoGen(hf_token)
	if err != nil {
		fmt.Println("error creating video: ", err)
	}

	fmt.Println("resp creating videos: ", videoPath)
}
