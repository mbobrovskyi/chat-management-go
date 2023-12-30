package http

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/internal/application/connector"
	"github.com/mbobrovskyi/chat-management-go/internal/data_contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/entities/user"
	service2 "github.com/mbobrovskyi/chat-management-go/internal/domain/services"
	"github.com/mbobrovskyi/chat-management-go/internal/infrastructure/server"
	"github.com/samber/lo"
)

type ChatController struct {
	authMiddleware server.Middleware
	chatService    service2.ChatService
	messageService service2.MessageService
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
	userSession := ctx.Context().UserValue("user").(*user.User)
	return websocket.New(func(conn *websocket.Conn) {
		websocketConnection := connection.NewConnection(conn, *userSession)
		c.chatConnector.AddConnection(ctx.Context(), websocketConnection)
		<-websocketConnection.GetCloseChan()
		fmt.Println("Connection closed on controller")
	})(ctx)
}

func (c *ChatController) getChats(ctx *fiber.Ctx) error {
	chats, count, err := c.chatService.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(data_contracts.NewPage[ChatResponse](lo.Map(chats, func(item chat.Chat, _ int) ChatResponse {
		return ChatToResponse(item)
	}), count))
}

func (c *ChatController) getChat(ctx *fiber.Ctx) error {
	return nil
}

func (c *ChatController) getChatMessages(ctx *fiber.Ctx) error {
	messages, count, err := c.messageService.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(data_contracts.NewPage[MessageResponse](lo.Map(messages, func(item message.Message, _ int) MessageResponse {
		return MessageToResponse(item)
	}), count))
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
	chatService service2.ChatService,
	messageService service2.MessageService,
	chatConnector connector.Connector,
) server.Controller {
	return &ChatController{
		authMiddleware: authMiddleware,
		chatService:    chatService,
		messageService: messageService,
		chatConnector:  chatConnector,
	}
}
