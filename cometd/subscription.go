package cometd

import "fmt"

type SubscriptionRequest struct {
	Channel      string `json:"channel"`
	Subscription string `json:"subscription"`
	Id           string `json:"id"`
	ClientId     string `json:"clientId"`
}

func (c *Client) Subscribe(channelId string) error {
	_, err := c.do(c.subscriptionRequest(channelId))
	return err
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
