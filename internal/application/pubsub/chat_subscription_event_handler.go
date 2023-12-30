package pubsub

import (
	"encoding/json"
	"github.com/mbobrovskyi/chat-management-go/internal/application/connector"
	"github.com/mbobrovskyi/chat-management-go/internal/common"
	domain2 "github.com/mbobrovskyi/chat-management-go/internal/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/services"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/logger"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/pubsub/subscriber"
)

type ChatSubscriptionHandler struct {
	log            logger.Logger
	messageService services.MessageService
	chatConnector  connector.Connector
}

func (c ChatSubscriptionHandler) Handle(eventType uint8, data []byte) error {
	switch eventType {
	case domain2.CreateMessagePubSubEventType:
		return c.createMessage(data)
	}

	return nil
}

func (c *ChatSubscriptionHandler) createMessage(data []byte) error {
	var dto common.MessageDTO

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	}

	for _, conn := range c.chatConnector.GetConnections() {
		if err := conn.SendEvent(domain2.CreateMessageWebsocketEventType, dto); err != nil {
			return err
		}
	}

	return nil
}

func NewChatSubscriberHandler(
	messageService services.MessageService,
	chatConnector connector.Connector,
) subscriber.EventHandler {
	return &ChatSubscriptionHandler{
		messageService: messageService,
		chatConnector:  chatConnector,
	}
}
