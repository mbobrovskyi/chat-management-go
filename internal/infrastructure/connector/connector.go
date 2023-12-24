package connector

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/domain/connection"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/logger"
	"sync"
	"time"
)

var AlreadyStartedError = errors.New("connector already started")

type Connector interface {
	Start(ctx context.Context) error
	AddConnection(conn connection.Connection)
	GetConnections() []connection.Connection
}

type connector struct {
	mtx           sync.RWMutex
	log           logger.Logger
	connections   []connection.Connection
	isStarted     bool
	eventHandler  connection.EventHandler
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

func (c *connector) AddConnection(conn connection.Connection) {
	c.log.Debugf("Added connection %d...", conn.GetId())

	conn.Connect()
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
			return
		case msg := <-conn.GetMessageChan():
			c.onMessage(conn, msg)
		}
	}
}

func (c *connector) onMessage(conn connection.Connection, data []byte) {
	var rawEvent connection.Event

	if err := json.Unmarshal(data, &rawEvent); err != nil {
		c.log.Debugf("Error on parse raw event: %s", err.Error())
		return
	}

	c.log.Debugf("Got new event event_type=%d message=%s", rawEvent.GetType(), string(rawEvent.GetData()))

	if err := c.eventHandler.HandleEvent(conn, rawEvent); err != nil {
		c.log.Error(err)
	}
}

func (c *connector) GetConnections() []connection.Connection {
	return c.connections
}

type Config struct {
	CleanInterval time.Duration
	Logger        logger.Logger
}

func NewConnector(
	eventHandler connection.EventHandler,
	configs ...Config,
) *connector {
	conn := &connector{
		log:           logger.NewNopLogger(),
		eventHandler:  eventHandler,
		cleanInterval: time.Minute,
	}

	for _, config := range configs {
		if config.CleanInterval > 0 {
			conn.cleanInterval = config.CleanInterval
		}

		if config.Logger != nil {
			conn.log = config.Logger
		}
	}

	return conn
}
