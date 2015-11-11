package yammer

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bmorton/go-yammer/schema"
)

type CreateMessageParams struct {
	Body        string `json:"body"`
	GroupId     int    `json:"group_id"`
	RepliedToId int    `json:"replied_to_id"`
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
