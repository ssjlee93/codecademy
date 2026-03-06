package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ssjlee93/recommendation-software/pkg/genai"
)

type chatRequest struct {
	Message string `json:"message"`
}

type chatResponse struct {
	Response string `json:"response"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// ChatHandler returns an http.HandlerFunc for POST /chat.
func ChatHandler(client *genai.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, errorResponse{"method not allowed"})
			return
		}

		var req chatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
			writeJSON(w, http.StatusBadRequest, errorResponse{"invalid request body: 'message' field is required"})
			return
		}

		msg := fmt.Sprintf("Recommend a Pokemon based on the user's request. Only respond with the Pokemon name. The user requested : %s", req.Message)
		reply, err := client.Chat(r.Context(), msg)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, errorResponse{"failed to get AI response"})
			return
		}

		writeJSON(w, http.StatusOK, chatResponse{Response: reply})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
