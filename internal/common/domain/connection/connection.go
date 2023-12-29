package connection

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/event"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/session"
)

var ClosedError = errors.New("connection closed")

type Connection interface {
	Equals(other Connection) bool

	GetUUID() uuid.UUID
	GetSession() session.Session
	GetMetadata() map[string]any

	IsOpened() bool
	IsClosed() bool

	GetMessageChan() chan []byte
	GetCloseChan() chan struct{}

	SendMessage(eventType uint8, data any) error

	Open()
	Close()
}

type connection struct {
	uuid     uuid.UUID
	metadata map[string]any

	conn    Conn
	session session.Session

	messageChan chan []byte
	closeChan   chan struct{}

	opened bool
	closed bool
}

func (c *connection) Equals(other Connection) bool {
	return c.GetUUID() == other.GetUUID()
}

func (c *connection) GetUUID() uuid.UUID {
	return c.uuid
}

func (c *connection) GetSession() session.Session {
	return c.session
}

func (c *connection) GetMetadata() map[string]any {
	return c.metadata
}

func (c *connection) IsOpened() bool {
	return c.opened
}

func (c *connection) IsClosed() bool {
	return c.closed
}

func (c *connection) GetMessageChan() chan []byte {
	return c.messageChan
}

func (c *connection) GetCloseChan() chan struct{} {
	return c.closeChan
}

func (c *connection) SendMessage(eventType uint8, data any) error {
	if c.closed {
		return ClosedError
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.conn.WriteJSON(event.Event{Type: eventType, Data: jsonData}); err != nil {
		return err
	}

	return nil
}

func (c *connection) Open() {
	go c.open()
}

func (c *connection) open() {
	if c.opened || c.closed {
		return
	}

	c.opened = true

	defer func() {
		c.opened = false
	}()

	for {
		select {
		case <-c.closeChan:
			return
		default:
			_, msgData, err := c.conn.ReadMessage()
			if err != nil {
				c.Close()
				return
			}
			c.messageChan <- msgData
		}
	}
}

func (c *connection) Close() {
	if !c.closed {
		close(c.closeChan)
		c.conn.Close()
	}

	c.closed = true
}
