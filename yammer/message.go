package yammer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"github.com/bmorton/go-yammer/schema"
	"github.com/google/go-querystring/query"
)

type CreateMessageParams struct {
	Body           string `json:"body"`
	GroupId        int    `json:"group_id"`
	RepliedToId    int    `json:"replied_to_id"`
	InvitedUserIDs string `json:"invited_user_ids"`
	AdditionalData string `json:"additional_data"`
}

func (c *Client) PostMessage(payload *CreateMessageParams) (*schema.Message, error) {
	resp, err := c.sendRequest(payload, "POST", "/api/v1/messages.json")
	if err != nil {
		return &schema.Message{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.Message{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.Message{}, err
	}

	return feed.Messages[0], nil
}

type MarkReadParams struct {
	ThreadID  int `url:"thread_id"`
	MessageID int `url:"message_id"`
}

var ErrMarkReadFailed = errors.New("Marking read did not return a successful response code")

func (c *Client) MarkRead(options *MarkReadParams) error {
	querystring, _ := query.Values(options)
	url := fmt.Sprintf("%s/api/v1/messages/last_seen_in_thread.json?%s", c.baseURL, querystring.Encode())

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	if resp.StatusCode != http.StatusOK {
		return ErrMarkReadFailed
	}

	return nil
}

type LikeMessageParams struct {
	MessageID int `url:"message_id"`
}

var ErrLikeMessageFailed = errors.New("Liking message did not return a successful response code")

func (c *Client) LikeMessage(options *LikeMessageParams) error {
	querystring, _ := query.Values(options)
	url := fmt.Sprintf("%s/api/v1/messages/liked_by/current.json?%s", c.baseURL, querystring.Encode())

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	if resp.StatusCode != http.StatusCreated {
		return ErrLikeMessageFailed
	}

	return nil
}
