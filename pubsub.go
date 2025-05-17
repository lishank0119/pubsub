package pubsub

import (
	"fmt"
	"hash/fnv"
)

type PubSub struct {
	buckets []*bucket
}

func (ps *PubSub) SubscriberCount(topic string) int {
	b := ps.getBucket(topic)
	return b.subscriberCount(topic)
}

func (ps *PubSub) ListTopics() []string {
	topics := make([]string, 0)
	for _, b := range ps.buckets {
		topics = append(topics, b.listTopics()...)
	}

	return topics
}

func (ps *PubSub) getBucket(topic string) *bucket {
	h := fnv.New32a()
	if _, err := h.Write([]byte(fmt.Sprintf("%v", topic))); err != nil {
		panic(fmt.Sprintf("unexpected error hashing key: %v", err))
	}
	i := int(h.Sum32()) % len(ps.buckets)
	return ps.buckets[i]
}

func (ps *PubSub) NewSubscriber() *Subscriber {
	return newSubscriber(ps)
}

// Publish a message to a topic
func (ps *PubSub) Publish(topic string, msg []byte) error {
	b := ps.getBucket(topic)
	return b.publish(&pubMessage{
		topic: topic,
		data:  msg,
	})
}

// UnsubscribeTopic all subscribers from a specific topic
func (ps *PubSub) UnsubscribeTopic(topic string) {
	b := ps.getBucket(topic)
	b.unsubscribeTopic(topic)
}

func (ps *PubSub) subscribe(s *Subscriber, topic string, handler HandlerFunc) {
	b := ps.getBucket(topic)
	b.subscribe(s, topic, handler)
}

func (ps *PubSub) unSubscribe(s *Subscriber, topic string) {
	b := ps.getBucket(topic)
	b.unSubscribe(s, topic)
}

func NewPubSub(config *Config) *PubSub {
	if config == nil {
		config = new(Config)
	}
	config.init()

	buckets := make([]*bucket, 0)

	for i := 0; i < config.BucketNum; i++ {
		buckets = append(buckets, newBucket(config.BucketMessageBuffer))
	}

	return &PubSub{
		buckets: buckets,
	}
}
