package cometd

import (
	"fmt"

	"github.com/bmorton/go-yammer/schema"
)

func (c *Client) Poll(messageChan chan *schema.MessageFeed, stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Aborting.")
			return
		default:
			resp, err := c.connect()
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, r := range resp {
				if r.Data != nil {
					messageChan <- r.Data.Data
				}
			}
		}
	}
}
