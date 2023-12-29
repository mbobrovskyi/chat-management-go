package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connector"
)

var InvalidConnectionType = errors.New("invalid connection type")

type ChatWebsocketEventHandler struct {
	messageService message.Service
}

func (h *ChatWebsocketEventHandler) Handle(baseConn connection.Connection, eventType uint8, data []byte) error {
	conn, ok := baseConn.(connection.Connection)
	if !ok {
		return InvalidConnectionType
	}

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
	var chatIds []uint64

	if err := json.Unmarshal(rawData, &chatIds); err != nil {
		return err
	}

	// TODO: Check chat ids on service is exist
	conn.GetMetadata()["currentChatIds"] = chatIds

	return nil
}

func (h *ChatWebsocketEventHandler) unsubscribeRoom(conn connection.Connection, _ []byte) error {
	delete(conn.GetMetadata(), "currentChatIds")
	return nil
}

func (h *ChatWebsocketEventHandler) setCurrentChat(conn connection.Connection, data []byte) error {
	var chatId uint64

	if err := json.Unmarshal(data, &chatId); err != nil {
		return err
	}

	// TODO: Check chat id on service is exist
	conn.GetMetadata()["currentChatId"] = chatId

	return nil
}

func (h *ChatWebsocketEventHandler) unsetCurrentChat(conn connection.Connection, _ []byte) error {
	delete(conn.GetMetadata(), "currentChatId")
	return nil
}

func (h *ChatWebsocketEventHandler) createMessage(conn connection.Connection, data []byte) error {
	request := MessageRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return err
	}

	currentChatId, ok := conn.GetMetadata()["currentChatId"]
	if !ok {
		return nil
	}

	msg := message.CreateNew(request.Text, currentChatId.(uint64), conn.GetSession().GetUser().GetId())

	if err := h.messageService.Create(context.Background(), msg); err != nil {
		return err
	}

	return nil
}

func (h *ChatWebsocketEventHandler) updateMessagesStatus(conn connection.Connection, data []byte) error {
	return nil
}

func NewMessageEventHandler(messageService message.Service) connector.EventHandler {
	return &ChatWebsocketEventHandler{messageService: messageService}
}
