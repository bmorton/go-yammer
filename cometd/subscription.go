package cometd

import (
	"fmt"
	"net/http"
)

type SubscriptionRequest struct {
	Channel      string `json:"channel"`
	Subscription string `json:"subscription"`
	Id           string `json:"id"`
	ClientId     string `json:"clientId"`
}

func (c *Client) SubscribeToChannel(channel string) error {
	_, err := c.do([]*SubscriptionRequest{
		&SubscriptionRequest{
			Channel:      "/meta/subscribe",
			Subscription: channel,
			Id:           c.nextRequestId(),
			ClientId:     c.clientId,
		},
	})

	return err
}

func (c *Client) SubscribeToFeed(channelId string) (*http.Response, error) {
	return c.do(c.subscriptionRequest(channelId))
}

func (c *Client) subscriptionRequest(channelId string) []*SubscriptionRequest {
	return []*SubscriptionRequest{
		&SubscriptionRequest{
			Channel:      "/meta/subscribe",
			Subscription: fmt.Sprintf("/feeds/%s/primary", channelId),
			Id:           c.nextRequestId(),
			ClientId:     c.clientId,
		},
		&SubscriptionRequest{
			Channel:      "/meta/subscribe",
			Subscription: fmt.Sprintf("/feeds/%s/secondary", channelId),
			Id:           c.nextRequestId(),
			ClientId:     c.clientId,
		},
	}
}
