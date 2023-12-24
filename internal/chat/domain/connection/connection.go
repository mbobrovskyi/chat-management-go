package connection

import (
	"encoding/json"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/aggregate_root"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/session"
)

type Connection interface {
	aggregate_root.AggregateRoot[Connection]

	IsClosed() bool
	GetSession() session.Session
	GetMessageChan() chan []byte
	GetCloseChan() chan struct{}
	SendEvent(eventType uint8, data any) error
	Connect()
	Close()
}

type connection struct {
	aggregate_root.AggregateRoot[Connection]

	conn    Conn
	session session.Session

	messageChan chan []byte
	closeChan   chan struct{}

	isConnected bool
	isClosed    bool
}

func (c *connection) IsClosed() bool {
	return c.isClosed
}

func (c *connection) GetSession() session.Session {
	return c.session
}

func (c *connection) GetMessageChan() chan []byte {
	return c.messageChan
}

func (c *connection) GetCloseChan() chan struct{} {
	return c.closeChan
}

func (c *connection) SendEvent(eventType uint8, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := c.conn.WriteJSON(NewEvent(eventType, jsonData)); err != nil {
		return err
	}

	return nil
}

func (c *connection) Connect() {
	go c.connect()
}

func (c *connection) connect() {
	if c.isConnected || c.isClosed {
		return
	}

	c.isConnected = true

	defer func() {
		c.isConnected = false
		c.isClosed = true
	}()

	for {
		select {
		case <-c.closeChan:
			return
		default:
			_, msgData, err := c.conn.ReadMessage()
			if err != nil {
				c.isConnected = false
				return
			}
			c.messageChan <- msgData
		}
	}
}

func (c *connection) Close() {
	if !c.isConnected {
		return
	}
	close(c.closeChan)
	c.conn.Close()
}
