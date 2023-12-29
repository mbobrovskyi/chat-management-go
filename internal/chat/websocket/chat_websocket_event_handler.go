package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/connection"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/connector"
)

var InvalidConnectionType = errors.New("invalid connection type")

type ChatWebsocketEventHandler struct {
	messageService domain.MessageService
}

func (h *ChatWebsocketEventHandler) Handle(baseConn connection.Connection, eventType uint8, data []byte) error {
	conn, ok := baseConn.(connection.Connection)
	if !ok {
		return InvalidConnectionType
	}

	fmt.Println("Handle:", eventType, string(data))

	switch eventType {
	case domain.SubscribeChatsWebsocketEventType:
		return h.subscribeChat(conn, data)
	case domain.UnsubscribeChatsWebsocketEventType:
		return h.unsubscribeRoom(conn, data)
	case domain.SetCurrentChatWebsocketEventType:
		return h.setCurrentChat(conn, data)
	case domain.UnsetCurrentChatWebsocketEventType:
		return h.unsetCurrentChat(conn, data)
	case domain.CreateMessageWebsocketEventType:
		return h.createMessage(conn, data)
	case domain.EditMessageWebsocketEventType:
	case domain.DeleteMessageWebsocketEventType:
	case domain.UpdateMessagesWebsocketStatusEventType:
		return h.updateMessagesStatus(conn, data)
	}

	return nil
}

func (h *ChatWebsocketEventHandler) subscribeChat(conn connection.Connection, rawData []byte) error {
	return nil
}

func (h *ChatWebsocketEventHandler) unsubscribeRoom(conn connection.Connection, rawData []byte) error {
	return nil
}

func (h *ChatWebsocketEventHandler) setCurrentChat(conn connection.Connection, rawData []byte) error {
	return nil
}

func (h *ChatWebsocketEventHandler) unsetCurrentChat(conn connection.Connection, rawData []byte) error {
	return nil
}

func (h *ChatWebsocketEventHandler) createMessage(conn connection.Connection, rawData []byte) error {
	dto := MessageDTO{}
	if err := json.Unmarshal(rawData, &dto); err != nil {
		return err
	}

	msg := message.CreateNew(dto.Text, 1, 1)

	if err := h.messageService.CreateMessage(context.Background(), msg); err != nil {
		return err
	}

	return nil
}

func (h *ChatWebsocketEventHandler) updateMessagesStatus(conn connection.Connection, rawData []byte) error {
	return nil
}

func NewMessageEventHandler(messageService domain.MessageService) connector.EventHandler {
	return &ChatWebsocketEventHandler{messageService: messageService}
}
