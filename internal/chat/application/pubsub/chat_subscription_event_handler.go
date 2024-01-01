package pubsub

import (
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/common/data_contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/connection"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/pubsub/subscriber"
	"github.com/mbobrovskyi/connector/pkg/connector"
)

type ChatSubscriptionHandler struct {
	log           logger.Logger
	chatConnector connector.Connector[*connection.Connection]
}

func (c *ChatSubscriptionHandler) Handle(eventType int, data []byte) error {
	switch eventType {
	case domain.CreateMessagePubSubEventType:
		return c.createMessage(data)
	}

	return nil
}

func (c *ChatSubscriptionHandler) createMessage(data []byte) error {
	var dto data_contracts.MessageDTO

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	for _, conn := range c.chatConnector.GetConnections() {
		if err := conn.SendEvent(domain.CreateMessageWebsocketEventType, dto); err != nil {
			return err
		}
	}

	return nil
}

func NewChatSubscriberHandler(
	chatConnector connector.Connector[*connection.Connection],
) subscriber.EventHandler {
	return &ChatSubscriptionHandler{
		chatConnector: chatConnector,
	}
}
