package main

import (
	"log"
	"os"
	"strconv"

	"github.com/bmorton/go-yammer/cometd"
	"github.com/bmorton/go-yammer/schema"
	"github.com/bmorton/go-yammer/yammer"
)

func main() {
	client := yammer.New(os.Getenv("YAMMER_TOKEN"))

	groupId, _ := strconv.Atoi(os.Getenv("YAMMER_GROUP"))
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

	rt.Subscribe(feed.Meta.Realtime.ChannelId)
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
