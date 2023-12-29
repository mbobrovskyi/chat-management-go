package connection

import (
	"github.com/google/uuid"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/session"
)

type Conn interface {
	WriteJSON(v any) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

func NewConnection(conn Conn, session session.Session) Connection {
	return &connection{
		uuid:        uuid.New(),
		metadata:    make(map[string]any),
		conn:        conn,
		session:     session,
		messageChan: make(chan []byte),
		closeChan:   make(chan struct{}),
	}
}
