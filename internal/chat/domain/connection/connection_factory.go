package connection

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/entity"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/session"
	"sync"
)

var lastId uint64 = 0
var lastIdMtx sync.Mutex

type Conn interface {
	WriteJSON(v any) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
}

func NewConnection(conn Conn, session session.Session) Connection {
	lastIdMtx.Lock()
	defer lastIdMtx.Unlock()

	lastId++

	return &connection{
		AggregateRoot: entity.New[Connection](lastId),

		conn:        conn,
		session:     session,
		messageChan: make(chan []byte),
		closeChan:   make(chan struct{}),
	}
}
