package api

import (
	"net/http"

	"github.com/ssjlee93/recommendation-software/pkg/genai"
)

// NewRouter creates and returns an http.ServeMux with all routes registered.
func NewRouter(client *genai.Client) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/chat", ChatHandler(client))
	return mux
}
