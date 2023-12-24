package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/infrastructure/server"
	"github.com/samber/lo"
)

var _ server.Controller = (*ChatController)(nil)

type ChatController struct {
	chatService domain.ChatService
}

func (c *ChatController) SetupRoutes(router fiber.Router) {
	chatGroup := router.Group("/chats")
	chatGroup.Get("", c.getChats)
	chatGroup.Get("/:id", c.getChat)
	chatGroup.Get("/:id/messages", c.getChatMessages)
	chatGroup.Post("", c.createChat)
	chatGroup.Put("/:id", c.updateChat)
	chatGroup.Delete("/:id", c.deleteChat)
}

func (c *ChatController) getChats(ctx *fiber.Ctx) error {
	chats, err := c.chatService.GetAll()
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

func NewChatController(chatService domain.ChatService) server.Controller {
	return &ChatController{chatService: chatService}
}
