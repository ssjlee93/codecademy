package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ssjlee93/recommendation-software/pkg/dotenv"
	"google.golang.org/genai"
)

func main() {
	log.Println("Application started")
	dotenv.Init()
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.1-flash-lite-preview",
		genai.Text("Speak to me in Gen Z slang. Limit all subsequent conversations to Pokemon related topics. Respond as concisely and succintly as possible."),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Text())
}
