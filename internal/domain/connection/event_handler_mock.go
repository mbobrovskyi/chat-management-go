package connection

type EventHandlerMock struct {
	handleEvent func(conn Connection, event Event) error
}

func (e *EventHandlerMock) HandleEvent(conn Connection, event Event) error {
	return e.handleEvent(conn, event)
}

func (e *EventHandlerMock) SetHandleEvent(handleEvent func(conn Connection, event Event) error) {
	e.handleEvent = handleEvent
}

func NewEventHandlerMock() *EventHandlerMock {
	return &EventHandlerMock{}
}
