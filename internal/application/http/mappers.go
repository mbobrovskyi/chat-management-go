package http

import (
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
)

func ChatToResponse(input chat.Chat) ChatResponse {
	return ChatResponse{
		Id:        input.Id(),
		Name:      input.Name(),
		Type:      input.Type().Uint8(),
		Image:     input.Image(),
		CreatedBy: input.CreatedBy(),
		CreatedAt: input.CreatedAt(),
		UpdatedAt: input.UpdatedAt(),
	}
}

func MessageToResponse(input message.Message) MessageResponse {
	return MessageResponse{
		Id:        input.Id(),
		Text:      input.Text(),
		Status:    input.Status().Uint8(),
		ChatId:    input.ChatId(),
		CreatedBy: input.CreatedById(),
		CreatedAt: input.CreatedAt(),
		UpdatedAt: input.UpdatedAt(),
	}
}
