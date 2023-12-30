package connector

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/connection"
)

type EventHandler interface {
	Handle(conn connection.Connection, eventType uint8, data []byte) error
}
