package cohere

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// RerankRequest represents a request structure for the rerank API.
// https://docs.cohere.com/reference/rerank-1
type RerankRequest struct {
	// Model defines the identifier of the model to use.
	Model string `json:"model"`

	// The search query
	Query string `json:"query"`

	// Documents is a list of document objects or strings to rerank.
	// If a document is provided the text fields is required and all other fields will be preserved in the response.
	// The total max chunks (length of documents * max_chunks_per_doc) must be less than 10000.
	Documents []string `json:"documents"`

	// TopN sets the number of most relevant documents or indices to return, defaults to the length of the documents
	TopN int `json:"top_n"`

	// ReturnDocuments:
	// * false: returns results without the doc text - the api will return a list of {index, relevance score},
	// where index is inferred from the list passed into the request.
	//
	// * true: returns results with the doc text passed in - the api will return an ordered list of
	// {index, text, relevance score} where index + text refers to the list passed into the request.
	ReturnDocuments bool `json:"return_documents"`

	// MaxChunksPerDoc sets the maximum number of chunks to produce internally from a document.
	MaxChunksPerDoc int `json:"max_chunks_per_doc,omitempty"`
}

type RerankResult struct {
	Index          int     `json:"index"`
	RelevanceScore float64 `json:"relevance_score"`
}

// RerankResponse is the response from the Cohere API for a rerank request.
type RerankResponse struct {
	ID      string         `json:"id"`
	Results []RerankResult `json:"results"`
}

func (c *Client) Rerank(ctx context.Context, req *RerankRequest) (*RerankResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/rerank", c.baseURL), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	response, err := c.DoRequest(request)
	if err != nil {
		return nil, err
	}

	var rsp RerankResponse
	err = json.NewDecoder(response.Body).Decode(&rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
