package pubsub

type HandlerFunc func(msg []byte)

type message struct {
	topic string
	data  []byte
}

type pubMessage struct {
	topic string
	data  []byte
}
