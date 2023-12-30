package connector

import (
	logger2 "github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"time"
)

type Config struct {
	CleanInterval time.Duration
	Logger        logger2.Logger
}

func NewConnector(
	eventHandler EventHandler,
	configs ...Config,
) Connector {
	conn := &connector{
		log:           logger2.NewNopLogger(),
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
