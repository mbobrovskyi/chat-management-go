package connector

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/event"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"sync"
	"time"
)

var AlreadyStartedError = errors.New("connector already started")

type Connector interface {
	Start(ctx context.Context) error
	AddConnection(ctx context.Context, conn connection.Connection)
	GetConnections() []connection.Connection
}

type connector struct {
	mtx           sync.RWMutex
	log           logger.Logger
	connections   []connection.Connection
	isStarted     bool
	eventHandler  EventHandler
	cleanInterval time.Duration
}

func (c *connector) Start(ctx context.Context) error {
	if c.isStarted {
		return AlreadyStartedError
	}

	c.isStarted = true
	defer func() {
		c.isStarted = false
	}()

	for {
		select {
		case <-ctx.Done():
			c.closeAll()
			return nil
		case <-time.After(c.cleanInterval):
			c.clean()
		}
	}
}

func (c *connector) closeAll() {
	c.log.Debug("Closing all connections...")

	c.mtx.Lock()
	defer c.mtx.Unlock()

	for _, conn := range c.connections {
		conn.Close()
	}
}

func (c *connector) clean() {
	c.log.Debug("Cleaning closed connections...")

	c.mtx.Lock()
	defer c.mtx.Unlock()

	connections := make([]connection.Connection, 0)

	for _, conn := range c.connections {
		if !conn.IsClosed() {
			connections = append(connections, conn)
		}
	}

	c.connections = connections
}

func (c *connector) AddConnection(ctx context.Context, conn connection.Connection) {
	c.log.Debugf("Added connection %s...", conn.GetUUID().String())

	conn.Open()
	c.addConnection(conn)
	go c.listen(conn)
}

func (c *connector) addConnection(conn connection.Connection) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.connections = append(c.connections, conn)
}

func (c *connector) listen(conn connection.Connection) {
	for {
		select {
		case <-conn.GetCloseChan():
			c.log.Debugf("Connection %s closed on connector", conn.GetUUID())
			return
		case msg := <-conn.GetMessageChan():
			c.onMessage(conn, msg)
		}
	}
}

func (c *connector) onMessage(conn connection.Connection, data []byte) {
	var rawEvent event.Event

	if err := json.Unmarshal(data, &rawEvent); err != nil {
		c.log.Debugf("Error on parse raw event: %s", err.Error())
		return
	}

	if err := c.eventHandler.Handle(conn, rawEvent.Type, rawEvent.Data); err != nil {
		c.log.Error(err)
	}
}

func (c *connector) GetConnections() []connection.Connection {
	return c.connections
}
