package subscriber

type EventHandler interface {
	Handle(eventType uint8, data []byte) error
}