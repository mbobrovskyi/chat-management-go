package application

import (
	"github.com/mbobrovskyi/ddd-chat-management-go/internal/chat/domain/chat"
)

func ChatToResponse(chat chat.Chat) ChatResponse {
	return ChatResponse{
		Id:        chat.GetId(),
		Name:      chat.GetName(),
		Type:      chat.GetType().Uint8(),
		Image:     chat.GetImage(),
		CreatedBy: chat.GetCreatedBy(),
		CreatedAt: chat.GetCreatedAt(),
		UpdatedAt: chat.GetUpdatedAt(),
	}
}
