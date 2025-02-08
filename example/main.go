package main

import (
	"fmt"
	"github.com/lishank0119/pubsub"
	"time"
)

func main() {
	ps := pubsub.NewPubSub()
	subscriber := ps.NewSubscriber(256)

	subscriber.Subscribe("news", func(msg []byte) {
		fmt.Println("Received:", string(msg))
	})

	ps.Publish("news", []byte("Hello, PubSub World!"))

	time.Sleep(1 * time.Second)
}
