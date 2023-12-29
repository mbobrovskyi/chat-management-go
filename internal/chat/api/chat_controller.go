package api

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/connector"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/session"
	"github.com/mbobrovskyi/chat-management-go/internal/common/domain/user"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/server"
	"github.com/samber/lo"
	"time"
)

var _ server.Controller = (*ChatController)(nil)

type ChatController struct {
	authMiddleware server.Middleware
	chatService    domain.ChatService
	chatConnector  connector.Connector
}

func (c *ChatController) SetupRoutes(router fiber.Router) {
	chatGroup := router.Group("/chats", c.authMiddleware.Handle)
	chatGroup.All("/ws", c.ws)
	chatGroup.Get("", c.getChats)
	chatGroup.Get("/:id", c.getChat)
	chatGroup.Get("/:id/messages", c.getChatMessages)
	chatGroup.Post("", c.createChat)
	chatGroup.Put("/:id", c.updateChat)
	chatGroup.Delete("/:id", c.deleteChat)
}

func (c *ChatController) ws(ctx *fiber.Ctx) error {
	currentSession := session.NewSession(user.NewUser(1, "test@test.com", "Test", "Test", time.Now(), time.Now()))
	return websocket.New(func(conn *websocket.Conn) {
		websocketConnection := connection.NewConnection(conn, currentSession)
		c.chatConnector.AddConnection(ctx.Context(), websocketConnection)
		<-websocketConnection.GetCloseChan()
		fmt.Println("Connection closed on controller")
	})(ctx)
}

func (c *ChatController) getChats(ctx *fiber.Ctx) error {
	chats, err := c.chatService.GetAll(ctx.Context())
	if err != nil {
		return err
	}
	return ctx.JSON(lo.Map(chats, func(chat chat.Chat, _ int) ChatResponse {
		return ChatToResponse(chat)
	}))
}

func (c *ChatController) getChat(ctx *fiber.Ctx) error {
	return nil
}

func (c *ChatController) getChatMessages(ctx *fiber.Ctx) error {
	return nil
}

func (c *ChatController) createChat(ctx *fiber.Ctx) error {
	return nil
}

func (c *ChatController) updateChat(ctx *fiber.Ctx) error {
	return nil
}

func (c *ChatController) deleteChat(ctx *fiber.Ctx) error {
	return nil
}

func NewChatController(
	authMiddleware server.Middleware,
	chatService domain.ChatService,
	chatConnector connector.Connector,
) server.Controller {
	return &ChatController{
		authMiddleware: authMiddleware,
		chatService:    chatService,
		chatConnector:  chatConnector,
	}
}
