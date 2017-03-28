package yammer

import (
	"fmt"

	"github.com/bmorton/go-yammer/schema"
)

type CreateActivityParams struct {
	Activities []*schema.Activity `json:"activity"`
}

func (c *Client) PostActivity(payload *CreateActivityParams) error {
	resp, err := c.sendRequest(payload, "POST", "/api/v1/activity.json")
	fmt.Println(resp)
	fmt.Println(err)

	return nil
}
