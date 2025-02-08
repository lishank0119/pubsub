package pubsub

import "sync"

type Subscriber struct {
	ps     *PubSub
	topics map[string]struct{}
	mu     sync.Mutex
}

func (s *Subscriber) addTopic(topic string) {
	s.mu.Lock()
	s.topics[topic] = struct{}{}
	s.mu.Unlock()
}

func (s *Subscriber) removeTopic(topic string) {
	s.mu.Lock()
	delete(s.topics, topic)
	s.mu.Unlock()
}

// Subscribe to a topic with a handler
func (s *Subscriber) Subscribe(topic string, handler HandlerFunc) {
	s.ps.subscribe(s, topic, handler)
}

// Unsubscribe from a specific topic
func (s *Subscriber) Unsubscribe(topic string) {
	s.ps.unSubscribe(s, topic)
}

// UnsubscribeAll from all topics
func (s *Subscriber) UnsubscribeAll() {
	var topics []string

	s.mu.Lock()
	for topic := range s.topics {
		topics = append(topics, topic)
	}
	s.mu.Unlock()

	for _, topic := range topics {
		s.ps.unSubscribe(s, topic)
	}
}

func newSubscriber(ps *PubSub) *Subscriber {
	s := &Subscriber{
		ps:     ps,
		topics: make(map[string]struct{}),
	}
	return s
}
