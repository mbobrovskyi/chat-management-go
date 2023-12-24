package connection

type EventHandler interface {
	HandleEvent(conn Connection, event Event) error
}
