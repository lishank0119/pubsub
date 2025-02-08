package pubsub

import (
	"errors"
	"sync"
)

var ErrMessageChannelFull = errors.New("message channel is full")

type bucket struct {
	mu     sync.RWMutex
	topics map[string]map[*Subscriber]HandlerFunc
	pubCh  chan *pubMessage
}

// unsubscribeTopic all subscribers from a specific topic
func (b *bucket) unsubscribeTopic(topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if subscribers, ok := b.topics[topic]; ok {
		for subscriber := range subscribers {
			subscriber.removeTopic(topic)
		}
	}

	delete(b.topics, topic)
}

func (b *bucket) unSubscribe(s *Subscriber, topic string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.topics[topic]; ok {
		delete(b.topics[topic], s)
		s.removeTopic(topic)
	}

	if len(b.topics[topic]) == 0 {
		delete(b.topics, topic)
	}
}

func (b *bucket) subscribe(s *Subscriber, topic string, handler HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.topics[topic]; !ok {
		b.topics[topic] = make(map[*Subscriber]HandlerFunc)
	}
	b.topics[topic][s] = handler
	s.addTopic(topic)
}

func (b *bucket) publish(msg *pubMessage) error {
	select {
	case b.pubCh <- msg:
	default:
		return ErrMessageChannelFull
	}
	return nil
}

func (b *bucket) start() {
	for msg := range b.pubCh {
		b.mu.RLock()
		if subscribers, ok := b.topics[msg.topic]; ok {
			for _, handler := range subscribers {
				handler(msg.data)
			}
		}
		b.mu.RUnlock()
	}
}

func newBucket(messageBuffer int) *bucket {
	if messageBuffer < 256 {
		messageBuffer = 256
	}

	b := &bucket{
		topics: make(map[string]map[*Subscriber]HandlerFunc),
		pubCh:  make(chan *pubMessage, messageBuffer),
	}

	go b.start()
	return b
}
