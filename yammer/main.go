package yammer

import "net/http"

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
