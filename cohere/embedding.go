package cohere

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// EmbeddingRequest represents a request structure for the /embed API.
// https://docs.cohere.com/reference/embed
type EmbeddingRequest struct {
	// Texts is an array of strings for the model to embed.
	// Maximum number of texts per call is 96.
	// We recommend reducing the length of each text to be under 512 tokens for optimal quality.
	Texts []string `json:"texts"`

	// Model is the identifier of the model.
	// Smaller "light" models are faster, while larger models will perform better.
	// Custom models can also be supplied with their full ID.
	// Defaults to embed-english-v2.0
	Model string `json:"model,omitempty"`

	// InputType specifies the type of input you're giving to the model.
	// Not required for older versions of the embedding models (i.e. anything lower than v3),
	// but is required for more recent versions (i.e. anything bigger than v2).
	InputType string `json:"input_type,omitempty"`

	// Truncate specifies how the API will handle inputs longer than the maximum token length.
	// One of NONE|START|END.
	// Passing START will discard the start of the input. END will discard the end of the input. In both cases, input is discarded until the remaining input is exactly the maximum input token length for the model.
	//
	//If NONE is selected, when the input exceeds the maximum input token length an error will be returned.
	//
	//Default: END
	Truncate string `json:"truncate,omitempty"`
}

// EmbeddingResponse is the response from the Cohere API for a embedding request.
type EmbeddingResponse struct {
	ID         string      `json:"id"`
	Texts      []string    `json:"texts"`
	Embeddings [][]float64 `json:"embeddings"`
}

func (c *Client) Embed(ctx context.Context, req *EmbeddingRequest) (*EmbeddingResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/embed", c.baseURL), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	response, err := c.DoRequest(request)
	if err != nil {
		return nil, err
	}

	var rsp EmbeddingResponse
	err = json.NewDecoder(response.Body).Decode(&rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
