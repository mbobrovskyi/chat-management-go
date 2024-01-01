package websocket

import (
	"context"
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/common/data_mappers"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/abstracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/publisher"
	"github.com/mbobrovskyi/connector/pkg/connector"
)

var _ connector.EventHandler[*connection.Connection] = (*ChatConnectorEventHandler)(nil)

type ChatConnectorEventHandler struct {
	chatRepository    abstracts.ChatRepository
	messageRepository abstracts.MessageRepository
	chatPublisher     publisher.Publisher
}

func (h *ChatConnectorEventHandler) Handle(conn *connection.Connection, eventType int, data []byte) error {
	switch eventType {
	case domain.SubscribeChatsWebsocketEventType:
		return h.subscribeChats(conn, data)
	case domain.UnsubscribeChatsWebsocketEventType:
		return h.unsubscribeChats(conn, data)
	case domain.OpenChatWebsocketEventType:
		return h.openChat(conn, data)
	case domain.CloseChatWebsocketEventType:
		return h.CloseChat(conn, data)
	case domain.CreateMessageWebsocketEventType:
		return h.createMessage(conn, data)
	}

	return nil
}

func (h *ChatConnectorEventHandler) subscribeChats(conn *connection.Connection, rawData []byte) error {
	var chatIds []uint64

	if err := json.Unmarshal(rawData, &chatIds); err != nil {
		return err
	}

	if err := conn.SubscribeChats(chatIds); err != nil {
		return err
	}

	return nil
}

func (h *ChatConnectorEventHandler) unsubscribeChats(conn *connection.Connection, _ []byte) error {
	conn.UnsubscribeChats()
	return nil
}

func (h *ChatConnectorEventHandler) openChat(conn *connection.Connection, data []byte) error {
	var chatId uint64

	if err := json.Unmarshal(data, &chatId); err != nil {
		return err
	}

	if err := conn.OpenChat(chatId); err != nil {
		return err
	}

	return nil
}

func (h *ChatConnectorEventHandler) CloseChat(conn *connection.Connection, _ []byte) error {
	conn.CloseChat()
	return nil
}

func (h *ChatConnectorEventHandler) createMessage(conn *connection.Connection, data []byte) error {
	request := CreateMessageRequest{}
	if err := json.Unmarshal(data, &request); err != nil {
		return err
	}

	if conn.CurrentChat() == nil {
		return nil
	}

	msg, err := message.Create(request.Text, *conn.CurrentChat(), *conn.User())
	if err != nil {
		return err
	}

	ctx := context.Background()

	newMsg, err := h.messageRepository.Save(ctx, msg)
	if err != nil {
		return err
	}

	if err := h.chatPublisher.Publish(ctx, domain.CreateMessagePubSubEventType, data_mappers.MessageToDTO(*newMsg)); err != nil {
		return err
	}

	return nil
}

func NewChatConnectorEventHandler(
	chatRepository abstracts.ChatRepository,
	messageRepository abstracts.MessageRepository,
	chatPublisher publisher.Publisher,
) *ChatConnectorEventHandler {
	return &ChatConnectorEventHandler{
		chatRepository:    chatRepository,
		messageRepository: messageRepository,
		chatPublisher:     chatPublisher,
	}
}
