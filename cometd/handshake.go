package cometd

import (
	"encoding/json"
	"io/ioutil"
)

type HandshakeRequest struct {
	Ext                      *Ext     `json:"ext"`
	Version                  string   `json:"version"`
	MinimumVersion           string   `json:"minimumVersion"`
	Channel                  string   `json:"channel"`
	SupportedConnectionTypes []string `json:"supportedConnectionTypes"`
	Id                       string   `json:"id"`
	ClientId                 string   `json:"clientId,omitempty"`
}

type Ext struct {
	Token string `json:"token"`
}

func (c *Client) Handshake() error {
	resp, err := c.do(c.handshakeRequest())
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var handshake []*HandshakeRequest
	err = json.Unmarshal(body, &handshake)
	if err != nil {
		return err
	}

	c.clientId = handshake[0].ClientId
	return nil
}

func (c *Client) handshakeRequest() []*HandshakeRequest {
	return []*HandshakeRequest{
		&HandshakeRequest{
			Ext: &Ext{
				Token: c.authenticationToken,
			},
			Id:                       c.nextRequestId(),
			Version:                  "1.0",
			MinimumVersion:           "0.9",
			Channel:                  "/meta/handshake",
			SupportedConnectionTypes: []string{"long-polling"},
		},
	}
}
