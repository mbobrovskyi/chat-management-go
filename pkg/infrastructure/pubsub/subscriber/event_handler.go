package subscriber

type EventHandler interface {
	Handle(eventType int, data []byte) error
}
