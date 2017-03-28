package yammer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bmorton/go-yammer/schema"
	"github.com/google/go-querystring/query"
)

type GetNoteOptions struct {
	IncludeBody          bool `url:"include_body,omitempty"`
	IncludeFeaturedUsers bool `url:"include_featured_users,omitempty"`
	IncludeVersions      bool `url:"include_versions,omitempty"`
}

func (c *Client) GetNote(noteId string, options GetNoteOptions) (*schema.Page, error) {
	querystring, _ := query.Values(options)

	url := fmt.Sprintf("/api/v1/notes/%s.json?%s", noteId, querystring.Encode())
	resp, err := c.sendRequest(nil, "GET", url)
	if err != nil {
		return &schema.Page{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.Page{}, err
	}

	var page schema.Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		return &schema.Page{}, err
	}

	return &page, nil
}
