package yammer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type Client struct {
	baseURL     string
	bearerToken string
	connection  *http.Client
	DebugMode   bool
}

func New(bearerToken string) *Client {
	return &Client{
		baseURL:     "https://www.yammer.com",
		bearerToken: bearerToken,
		connection:  &http.Client{},
	}
}

func (c *Client) Token() string {
	return c.bearerToken
}

func (c *Client) SetBaseURL(url string) {
	c.baseURL = url
}

func (c *Client) SetStaging() {
	c.SetBaseURL("https://www.staging.yammer.com")
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
	req.Header.Add("Yammer-Capabilities", "external_messaging,external_groups,system_request,user_sidebar,parsed_body_only2")

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	return resp, err
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
