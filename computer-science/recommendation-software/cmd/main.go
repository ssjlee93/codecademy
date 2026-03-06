package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ssjlee93/recommendation-software/pkg/api"
	"github.com/ssjlee93/recommendation-software/pkg/dotenv"
	genaiClient "github.com/ssjlee93/recommendation-software/pkg/genai"
)

func main() {
	log.Println("Starting server...")
	dotenv.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	ctx := context.Background()
	client, err := genaiClient.New(ctx, os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatalf("failed to initialize GenAI client: %v", err)
	}

	mux := api.NewRouter(client)

	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
