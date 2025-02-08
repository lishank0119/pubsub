package pubsub

import (
	"sync"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	ps := NewPubSub()
	subscriber := ps.NewSubscriber(256)

	var wg sync.WaitGroup
	wg.Add(1)

	// Test Subscribe and Publish
	received := false
	subscriber.Subscribe("topic1", func(msg []byte) {
		if string(msg) == "Hello, Subscriber!" {
			received = true
			wg.Done()
		}
	})

	ps.Publish("topic1", []byte("Hello, Subscriber!"))

	_ = waitWithTimeout(&wg, 1*time.Second)

	if !received {
		t.Errorf("Expected to receive message, but did not")
	}

	// Test Unsubscribe
	subscriber.Unsubscribe("topic1")
	received = false
	wg.Add(1)

	ps.Publish("topic1", []byte("This should not be received"))

	if waitWithTimeout(&wg, 500*time.Millisecond) == nil {
		t.Errorf("Expected not to receive message after unsubscribe")
	}

	// Test UnsubscribeAll
	subscriber.Subscribe("topic2", func(msg []byte) {
		received = true
		wg.Done()
	})
	subscriber.UnsubscribeAll()

	received = false
	wg.Add(1)
	ps.Publish("topic2", []byte("This should not be received"))

	if waitWithTimeout(&wg, 500*time.Millisecond) == nil {
		t.Errorf("Expected not to receive message after UnsubscribeAll")
	}

	// Test PubSub.UnsubscribeTopic
	subscriber.Subscribe("topic3", func(msg []byte) {
		received = true
		wg.Done()
	})

	ps.UnsubscribeTopic("topic3")

	received = false
	wg.Add(1)
	ps.Publish("topic3", []byte("This should not be received"))

	if waitWithTimeout(&wg, 500*time.Millisecond) == nil {
		t.Errorf("Expected not to receive message after PubSub.UnsubscribeTopic")
	}

	// Test PubSub.UnsubscribeAll
	subscriber.Subscribe("topic4", func(msg []byte) {
		received = true
		wg.Done()
	})

	ps.UnsubscribeAll()

	received = false
	wg.Add(1)
	ps.Publish("topic4", []byte("This should not be received"))

	if waitWithTimeout(&wg, 500*time.Millisecond) == nil {
		t.Errorf("Expected not to receive message after PubSub.UnsubscribeAll")
	}
}

func waitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		return nil
	case <-time.After(timeout):
		return &timeoutError{}
	}
}

type timeoutError struct{}

func (e *timeoutError) Error() string {
	return "timeout occurred"
}
