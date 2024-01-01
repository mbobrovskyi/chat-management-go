package connection

import (
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/user"
	"github.com/mbobrovskyi/connector/pkg/connector/connection/websocket"
)

type Connection struct {
	*websocket.Connection

	user            *user.User
	subscribedChats []uint64
	openedChat      *uint64
}

func (c *Connection) User() *user.User {
	return c.user
}

func (c *Connection) SubscribedChats() []uint64 {
	return c.subscribedChats
}

func (c *Connection) SubscribeChats(chats []uint64) error {
	// TODO: Check is exist by interface
	c.subscribedChats = chats
	return nil
}

func (c *Connection) UnsubscribeChats() {
	c.subscribedChats = nil
}

func (c *Connection) CurrentChat() *uint64 {
	return c.openedChat
}

func (c *Connection) OpenChat(chatId uint64) error {
	// TODO: Check is exist for current user
	c.openedChat = &chatId
	return nil
}

func (c *Connection) CloseChat() {
	c.openedChat = nil
}

func NewConnection(conn websocket.Conn, user *user.User) *Connection {
	return &Connection{
		Connection: websocket.New(conn),
		user:       user,
	}
}
