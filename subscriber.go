package pubsub

import "sync"

type Subscriber struct {
	ps       *PubSub
	handlers map[string]HandlerFunc
	messages chan message
	mu       sync.Mutex
}

func (s *Subscriber) sendMessage(msg message) {
	select {
	case s.messages <- msg:
	default:
	}
}

func (s *Subscriber) start() {
	for msg := range s.messages {
		s.mu.Lock()
		if handler, ok := s.handlers[msg.topic]; ok {
			handler(msg.data)
		}
		s.mu.Unlock()
	}
}

// Subscribe to a topic with a handler
func (s *Subscriber) Subscribe(topic string, handler HandlerFunc) {
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()

	s.mu.Lock()
	s.handlers[topic] = handler
	s.mu.Unlock()

	if s.ps.topics[topic] == nil {
		s.ps.topics[topic] = make(map[*Subscriber]struct{})
	}
	s.ps.topics[topic][s] = struct{}{}
}

// Unsubscribe from a specific topic
func (s *Subscriber) Unsubscribe(topic string) {
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.handlers[topic]; ok {
		delete(s.handlers, topic)
		delete(s.ps.topics[topic], s)

		if len(s.ps.topics[topic]) == 0 {
			delete(s.ps.topics, topic)
		}
	}
}

// UnsubscribeAll from all topics
func (s *Subscriber) UnsubscribeAll() {
	s.ps.mu.Lock()
	defer s.ps.mu.Unlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	for topic := range s.handlers {
		delete(s.ps.topics[topic], s)
		if len(s.ps.topics[topic]) == 0 {
			delete(s.ps.topics, topic)
		}
	}
	s.handlers = make(map[string]HandlerFunc)
}

func newSubscriber(ps *PubSub, messageBuffer int) *Subscriber {
	if messageBuffer < 256 {
		messageBuffer = 256
	}

	s := &Subscriber{
		ps:       ps,
		handlers: make(map[string]HandlerFunc),
		messages: make(chan message, messageBuffer),
	}
	go s.start()
	return s
}
