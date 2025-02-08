package pubsub

import (
	"sync"
)

type PubSub struct {
	mu     sync.RWMutex
	topics map[string]map[*Subscriber]struct{}
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string]map[*Subscriber]struct{}),
	}
}

// NewSubscriber creates a new Subscriber associated with the given PubSub instance.
// The messageBuffer parameter defines the buffer size for the subscriber's message channel.
// If the provided messageBuffer is less than 256, it defaults to 256 to ensure efficient handling.
func (ps *PubSub) NewSubscriber(messageBuffer int) *Subscriber {
	return newSubscriber(ps, messageBuffer)
}

// Publish a message to a topic
func (ps *PubSub) Publish(topic string, msg []byte) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if subs, ok := ps.topics[topic]; ok {
		for sub := range subs {
			sub.sendMessage(message{topic: topic, data: msg})
		}
	}
}

// UnsubscribeTopic all subscribers from a specific topic
func (ps *PubSub) UnsubscribeTopic(topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if subs, ok := ps.topics[topic]; ok {
		for sub := range subs {
			sub.mu.Lock()
			delete(sub.handlers, topic)
			sub.mu.Unlock()
		}
		delete(ps.topics, topic)
	}
}

// UnsubscribeAll all subscribers from all topics
func (ps *PubSub) UnsubscribeAll() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for topic, subs := range ps.topics {
		for sub := range subs {
			sub.mu.Lock()
			delete(sub.handlers, topic)
			sub.mu.Unlock()
		}
		delete(ps.topics, topic)
	}
}
