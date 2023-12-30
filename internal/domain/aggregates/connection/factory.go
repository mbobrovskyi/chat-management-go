package connection

import (
	"github.com/google/uuid"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
)

type Conn interface {
	WriteJSON(v any) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

func NewConnection(conn Conn, user user.User) Connection {
	return &connection{
		uuid:        uuid.New(),
		metadata:    make(map[string]any),
		conn:        conn,
		user:        user,
		messageChan: make(chan []byte),
		closeChan:   make(chan struct{}),
	}
}
