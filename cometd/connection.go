package cometd

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bmorton/go-yammer/schema"
)

type ConnectionRequest struct {
	Channel        string `json:"channel"`
	ConnectionType string `json:"connectionType"`
	Id             string `json:"id"`
	ClientId       string `json:"clientId"`
}

type ConnectionResponse struct {
	Channel string          `json:"channel"`
	Data    *ConnectionData `json:"data"`
}

type ConnectionData struct {
	Data *schema.MessageFeed `json:"data"`
	Type string              `json:"type"`
}

func (c *Client) connect() ([]*ConnectionResponse, error) {
	resp, err := c.do(c.connectionRequest())
	if err != nil {
		return []*ConnectionResponse{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []*ConnectionResponse{}, err
	}

	var connResp []*ConnectionResponse
	err = json.Unmarshal(body, &connResp)
	if err != nil {
		return []*ConnectionResponse{}, err
	}

	return connResp, nil
}

func (c *Client) connectionRequest() []*ConnectionRequest {
	return []*ConnectionRequest{
		&ConnectionRequest{
			Id:             c.nextRequestId(),
			ClientId:       c.clientId,
			Channel:        "/meta/connect",
			ConnectionType: "long-polling",
		},
	}
}
