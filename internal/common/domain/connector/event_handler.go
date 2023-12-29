package connector

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/connection"
)

type EventHandler interface {
	Handle(conn connection.Connection, eventType uint8, data []byte) error
}
