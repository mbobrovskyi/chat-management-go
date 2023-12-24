package application

import (
	"fmt"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/connection"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/connector"
)

var _ connector.EventHandler = (*ChatEventHandler)(nil)

type ChatEventHandler struct{}

func (e *ChatEventHandler) Handle(conn connection.Connection, event connection.Event) error {
	fmt.Println(conn, event)
	return nil
}

func NewChatEventHandler() *ChatEventHandler {
	return &ChatEventHandler{}
}
