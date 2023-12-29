package subscription

import (
	"encoding/json"
	"fmt"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/common"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connector"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/subscriber"
)

type ChatPubSubHandler struct {
	messageService message.Service
	chatConnector  connector.Connector
}

func (c ChatPubSubHandler) Handle(eventType uint8, data []byte) error {
	fmt.Println("Handle:", eventType, string(data))

	switch eventType {
	case domain.CreateMessagePubSubEventType:
		return c.createMessage(data)
	}

	return nil
}

func (c *ChatPubSubHandler) createMessage(data []byte) error {
	var dto common.MessageDTO

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
	messageService message.Service,
	chatConnector connector.Connector,
) subscriber.EventHandler {
	return &ChatPubSubHandler{
		messageService: messageService,
		chatConnector:  chatConnector,
	}
}
