package genai

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

const model = "gemini-3.1-flash-lite-preview"

// Client wraps the genai client.
type Client struct {
	inner *genai.Client
}

// New initializes and returns a new Client.
func New(ctx context.Context, apiKey string) (*Client, error) {
	inner, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("genai.New: %w", err)
	}
	return &Client{inner: inner}, nil
}

// Chat sends a message to the GenAI model and returns the text response.
func (c *Client) Chat(ctx context.Context, message string) (string, error) {
	result, err := c.inner.Models.GenerateContent(
		ctx,
		model,
		genai.Text(message),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("genai.Chat: %w", err)
	}
	return result.Text(), nil
}
