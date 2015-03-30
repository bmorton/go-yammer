# go-yammer

This is a very early, WIP Golang implementation of the essentials required to post messages to a Yammer network and monitoring a realtime stream for updates.  Don't depend on this library for anything important yet.

## Example

```go
package main

import (
	"log"

	"github.com/bmorton/go-yammer/cometd"
	"github.com/bmorton/go-yammer/schema"
	"github.com/bmorton/go-yammer/yammer"
)

func main() {
	client := yammer.New("your-token-here")
	groupId := 12345

	feed, err := client.GroupFeed(groupId)
	if err != nil {
		log.Println(err)
		return
	}

	rt := cometd.New(feed.Meta.Realtime.URI, feed.Meta.Realtime.AuthenticationToken)
	err = rt.Handshake()
	if err != nil {
		log.Println(err)
		return
	}

	err = rt.Subscribe(feed.Meta.Realtime.ChannelId)
	if err != nil {
		log.Println(err)
		return
	}

	messageChan := make(chan *schema.MessageFeed, 10)
	stopChan := make(chan bool, 1)

	log.Printf("Polling group ID: %d\n", groupId)
	go rt.Poll(messageChan, stopChan)
	for {
		select {
		case m := <-messageChan:
			message := m.Messages[0]
			if message.IsThreadStarter() {
				log.Printf("[%d] %s", message.SenderId, message.Body.Plain)
			}
		}
	}
}
```