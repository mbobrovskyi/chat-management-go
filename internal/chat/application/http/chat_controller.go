package http

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/abstracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/connection"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/message"
	"github.com/mbobrovskyi/chat-management-go/pkg/application/http"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/baseerror"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/entities/user"
	"github.com/mbobrovskyi/chat-management-go/pkg/domain/errors"
	"github.com/mbobrovskyi/chat-management-go/pkg/infrastructure/httpserver"
	"github.com/mbobrovskyi/connector/pkg/connector"
	"github.com/samber/lo"
	"strconv"
)

type ChatController struct {
	authMiddleware    httpserver.Middleware
	chatRepository    abstracts.ChatRepository
	messageRepository abstracts.MessageRepository
	chatConnector     connector.Connector[*connection.Connection]
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
	user := ctx.Context().UserValue("user").(*user.User)
	return websocket.New(func(websocketConn *websocket.Conn) {
		conn := connection.NewConnection(websocketConn, user)
		c.chatConnector.AddConnection(conn)
		<-conn.CloseChan()
		fmt.Println("Connection closed on controller")
	})(ctx)
}

func (c *ChatController) getChats(ctx *fiber.Ctx) error {
	chats, count, err := c.chatRepository.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.NewPage[ChatResponse](lo.Map(chats, func(item chat.Chat, _ int) ChatResponse {
		return ChatToResponse(item)
	}), count))
}

func (c *ChatController) getChat(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return errors.NewValueIsNotValidError().WithMetadata("id", idStr)
	}

	chat, err := c.chatRepository.GetById(ctx.Context(), id)
	if err != nil {
		return err
	}

	if chat == nil {
		return baseerror.NewNotFoundError("Chat not found.")
	}

	return ctx.JSON(ChatToResponse(*chat))
}

func (c *ChatController) getChatMessages(ctx *fiber.Ctx) error {
	messages, count, err := c.messageRepository.GetAll(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.NewPage[MessageResponse](lo.Map(messages, func(item message.Message, _ int) MessageResponse {
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
	authMiddleware httpserver.Middleware,
	chatRepository abstracts.ChatRepository,
	messageRepository abstracts.MessageRepository,
	chatConnector connector.Connector[*connection.Connection],
) httpserver.Controller {
	return &ChatController{
		authMiddleware:    authMiddleware,
		chatRepository:    chatRepository,
		messageRepository: messageRepository,
		chatConnector:     chatConnector,
	}
}
