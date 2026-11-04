package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// StreamResponse represents a chunk of the streaming response
type StreamResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   *Usage   `json:"usage,omitempty"`
}

// Choice represents a single choice in the response
type Choice struct {
	Index        int     `json:"index"`
	Delta        Delta   `json:"delta"`
	FinishReason *string `json:"finish_reason"`
}

// Delta represents the incremental content
type Delta struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// Usage represents token usage statistics
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Request represents the incoming request
type Request struct {
	Model    string    `json:"model,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	Stream   bool      `json:"stream,omitempty"`
}

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func handleResponses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set headers for Server-Sent Events
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Mock response text to stream
	mockResponse := "Hello! This is a mock streaming response from the OpenAI-compatible API. " +
		"The server is working correctly and streaming data chunk by chunk. " +
		"You can customize this message or make it dynamic based on the request."

	// Generate a unique ID for this response
	responseID := fmt.Sprintf("chatcmpl-%d", time.Now().Unix())
	timestamp := time.Now().Unix()

	// Calculate mock token counts
	promptTokens := estimateTokens(req.Messages)
	completionTokens := estimateTokens([]Message{{Content: mockResponse}})

	// Send initial chunk with role
	initialChunk := StreamResponse{
		ID:      responseID,
		Object:  "chat.completion.chunk",
		Created: timestamp,
		Model:   "gpt-3.5-turbo",
		Choices: []Choice{
			{
				Index: 0,
				Delta: Delta{
					Role: "assistant",
				},
				FinishReason: nil,
			},
		},
	}

	data, _ := json.Marshal(initialChunk)
	fmt.Fprintf(w, "data: %s\n\n", data)
	flusher.Flush()

	// Stream the response token by token
	words := splitIntoTokens(mockResponse)
	for _, word := range words {
		chunk := StreamResponse{
			ID:      responseID,
			Object:  "chat.completion.chunk",
			Created: timestamp,
			Model:   "gpt-3.5-turbo",
			Choices: []Choice{
				{
					Index: 0,
					Delta: Delta{
						Content: word,
					},
					FinishReason: nil,
				},
			},
		}

		data, _ := json.Marshal(chunk)
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()

		// Simulate streaming delay
		time.Sleep(50 * time.Millisecond)
	}

	// Send final chunk with finish_reason and usage
	finishReason := "stop"
	finalChunk := StreamResponse{
		ID:      responseID,
		Object:  "chat.completion.chunk",
		Created: timestamp,
		Model:   "gpt-3.5-turbo",
		Choices: []Choice{
			{
				Index:        0,
				Delta:        Delta{},
				FinishReason: &finishReason,
			},
		},
		Usage: &Usage{
			PromptTokens:     promptTokens,
			CompletionTokens: completionTokens,
			TotalTokens:      promptTokens + completionTokens,
		},
	}

	data, _ = json.Marshal(finalChunk)
	fmt.Fprintf(w, "data: %s\n\n", data)
	fmt.Fprintf(w, "data: [DONE]\n\n")
	flusher.Flush()
}

// estimateTokens provides a rough estimate of token count
// Real tokenization is more complex, but this gives a reasonable approximation
func estimateTokens(messages []Message) int {
	totalWords := 0
	for _, msg := range messages {
		words := strings.Fields(msg.Content)
		totalWords += len(words)
	}
	// Rough estimate: ~1.3 tokens per word on average
	return int(float64(totalWords) * 1.3)
}

// splitIntoTokens splits text into words/tokens for streaming
func splitIntoTokens(text string) []string {
	var tokens []string
	currentToken := ""

	for _, char := range text {
		currentToken += string(char)
		if char == ' ' || char == '.' || char == ',' || char == '!' || char == '?' {
			tokens = append(tokens, currentToken)
			currentToken = ""
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func main() {
	http.HandleFunc("/v1/responses", handleResponses)

	port := "8090"
	log.Printf("Starting OpenAI-compatible mock server on port %s", port)
	log.Printf("Endpoint: http://localhost:%s/v1/responses", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
