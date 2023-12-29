package subscription

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/connector"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/common/domain/pubsub"
)

type ChatPubSubHandler struct {
	messageService domain.MessageService
	chatConnector  connector.Connector
}

func (c ChatPubSubHandler) Handle(eventType uint8, data []byte) error {
	switch eventType {
	case domain.CreateMessagePubSubEventType:
		return c.createMessage(data)
	}

	return nil
}

func (c *ChatPubSubHandler) createMessage(data []byte) error {
	for _, conn := range c.chatConnector.GetConnections() {
		if err := conn.SendMessage(domain.CreateMessageWebsocketEventType, data); err != nil {
			return err
		}
	}
	return nil
}

func NewChatSubscriberHandler(
	messageService domain.MessageService,
	chatConnector connector.Connector,
) pubsub.EventHandler {
	return &ChatPubSubHandler{
		messageService: messageService,
		chatConnector:  chatConnector,
	}
}
