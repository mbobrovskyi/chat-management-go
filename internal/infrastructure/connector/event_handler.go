package connector

import "github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/connection"

type EventHandler interface {
	Handle(conn connection.Connection, event connection.Event) error
}
