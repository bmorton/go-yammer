package yammer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bmorton/go-yammer/schema"
	"github.com/google/go-querystring/query"
)

func (c *Client) VisitGroup(groupId string) error {
	resp, err := c.sendRequest(nil, "POST", fmt.Sprintf("/api/v1/groups/%s/visit", groupId))

	if resp.StatusCode != http.StatusCreated {
		return errors.New("a valid status code was not returned")
	}

	return err
}

type GetGroupOptions struct {
	IncludeGroupInfo       bool `url:"include_group_info,omitempty"`
	IncludeRelatedContents bool `url:"include_related_contents,omitempty"`
}

func (c *Client) GetGroup(groupId string, options GetGroupOptions) (*schema.Group, error) {
	querystring, _ := query.Values(options)
	url := fmt.Sprintf("/api/v1/groups/%s.json?%s", groupId, querystring.Encode())
	resp, err := c.sendRequest(nil, "GET", url)
	if err != nil {
		return &schema.Group{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.Group{}, err
	}

	var group schema.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return &schema.Group{}, err
	}

	return &group, nil
}

func (c *Client) GetGroupNotes(groupId string) (*schema.PageCollection, error) {
	url := fmt.Sprintf("/api/v1/notes/in_group/%s.json", groupId)
	resp, err := c.sendRequest(nil, "GET", url)
	if err != nil {
		return &schema.PageCollection{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.PageCollection{}, err
	}

	var pages schema.PageCollection
	err = json.Unmarshal(body, &pages)
	if err != nil {
		return &schema.PageCollection{}, err
	}

	return &pages, nil
}
