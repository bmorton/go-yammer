package cometd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Client struct {
	uri                 string
	authenticationToken string
	connection          *http.Client
	globalRequestId     int
	clientId            string
	counterMutex        *sync.Mutex
}

func New(uri string, authenticationToken string) *Client {
	return &Client{
		uri:                 uri,
		authenticationToken: authenticationToken,
		connection:          &http.Client{},
		globalRequestId:     0,
		counterMutex:        &sync.Mutex{},
	}
}

func (c *Client) nextRequestId() string {
	c.counterMutex.Lock()
	nextId := c.globalRequestId + 1
	c.globalRequestId = nextId
	c.counterMutex.Unlock()
	return strconv.Itoa(nextId)
}

func (c *Client) do(payload interface{}) (*http.Response, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(payload)

	req, err := http.NewRequest("POST", c.uri, &b)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	return c.connection.Do(req)
}
