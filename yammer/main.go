package yammer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL     string
	bearerToken string
	connection  *http.Client
}

func New(bearerToken string) *Client {
	return &Client{
		baseURL:     "https://www.yammer.com",
		bearerToken: bearerToken,
		connection:  &http.Client{},
	}
}

func (c *Client) SetStaging() {
	c.baseURL = "https://www.staging.yammer.com"
}

func (c *Client) sendRequest(payload interface{}, verb string, endpoint string) (*http.Response, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(payload)

	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	req, err := http.NewRequest(verb, url, &b)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Add("Content-Type", "application/json")

	return c.connection.Do(req)
}
