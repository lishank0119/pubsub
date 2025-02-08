package main

import (
	"fmt"
	"github.com/lishank0119/pubsub"
	"time"
)

func main() {
	ps := pubsub.NewPubSub(nil)
	subscriber := ps.NewSubscriber()

	go Publish(ps)

	subscriber.Subscribe("news", func(msg []byte) {
		fmt.Println("Received:", string(msg))
	})

	subscriber.Subscribe("news:2", func(msg []byte) {
		fmt.Println("Received(2):", string(msg))
	})

	time.Sleep(1500 * time.Millisecond)

	subscriber.Unsubscribe("news:2")

	time.Sleep(1 * time.Second)

	subscriber.UnsubscribeAll()

	select {}
}

func Publish(ps *pubsub.PubSub) {
	IntervalTime := 1 * time.Second
	ticker := time.NewTicker(IntervalTime)
	for {
		select {
		case <-ticker.C:
			if err := ps.Publish("news", []byte("Hello, PubSub World!")); err != nil {
				panic(err)
				return
			}

			if err := ps.Publish("news:2", []byte("Hello, PubSub World!(2)")); err != nil {
				panic(err)
				return
			}
		}
	}
}
