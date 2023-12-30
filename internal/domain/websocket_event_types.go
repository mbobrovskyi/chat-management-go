package domain

const (
	SubscribeChatsWebsocketEventType       uint8 = 1
	UnsubscribeChatsWebsocketEventType           = 2
	SetCurrentChatWebsocketEventType             = 3
	UnsetCurrentChatWebsocketEventType           = 4
	CreateMessageWebsocketEventType              = 5
	EditMessageWebsocketEventType                = 6
	DeleteMessageWebsocketEventType              = 7
	UpdateMessagesWebsocketStatusEventType       = 8
)
