package yammer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmorton/go-yammer/schema"
)

type CreateMessageParams struct {
	Body        string `json:"body"`
	GroupId     int    `json:"group_id"`
	RepliedToId int    `json:"replied_to_id"`
}

func (c *Client) PostMessage(payload *CreateMessageParams) (*schema.Message, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(payload)

	url := fmt.Sprintf("%s/api/v1/messages.json", c.baseURL)
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return &schema.Message{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
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
