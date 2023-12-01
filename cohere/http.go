package cohere

import (
	"errors"
	"fmt"
	"net/http"
)

// DoRequest sends an HTTP request and returns the response, handling any non-OK HTTP status codes.
func (c *Client) DoRequest(request *http.Request) (*http.Response, error) {
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")
	request.Header.Set("authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("status code: %d", response.StatusCode))
	}

	return response, nil
}
