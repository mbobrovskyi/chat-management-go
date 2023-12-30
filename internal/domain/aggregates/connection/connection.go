package connection

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/valueobjects"
)

type Connection interface {
	Equals(other Connection) bool

	GetUUID() uuid.UUID
	GetUser() user.User
	GetMetadata() map[string]any

	IsOpened() bool
	IsClosed() bool

	GetMessageChan() chan []byte
	GetCloseChan() chan struct{}

	SendEvent(eventType uint8, data any) error

	Open()
	Close()
}

type connection struct {
	uuid     uuid.UUID
	metadata map[string]any

	conn Conn
	user user.User

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

func (c *connection) GetUser() user.User {
	return c.user
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

func (c *connection) SendEvent(eventType uint8, data any) error {
	if c.closed {
		return nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.conn.WriteJSON(valueobjects.Event{Type: eventType, Data: jsonData}); err != nil {
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
