package yammer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmorton/go-yammer/schema"
)

func (c *Client) GroupFeed(id int) (*schema.MessageFeed, error) {
	url := fmt.Sprintf("%s/api/v1/messages/in_group/%d.json", c.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.MessageFeed{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	return &feed, nil
}

func (c *Client) InboxFeed() (*schema.MessageFeed, error) {
	url := fmt.Sprintf("%s/api/v1/messages/inbox.json", c.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.MessageFeed{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	return &feed, nil
}
