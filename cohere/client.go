package cohere

import (
	"errors"
	"net/http"
	"time"
)

const cohereAPIURL = "https://api.cohere.ai/v1"

// Client represents the Cohere API client and its config.
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewClient initializes a new Cohere API client with the required headers.
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("missing API key")
	}

	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Minute * 5,
		},
		apiKey:  apiKey,
		baseURL: cohereAPIURL,
	}
	return client, nil
}
