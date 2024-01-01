package http

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/chat"
	"github.com/mbobrovskyi/chat-management-go/pkg/baseconfig"
)

func ChatToResponse(input chat.Chat) ChatResponse {
	return ChatResponse{
		Id:        input.Id(),
		Name:      input.Name(),
		Type:      input.Type().Uint8(),
		Image:     input.Image(),
		CreatedBy: input.CreatedBy(),
		CreatedAt: input.CreatedAt().Format(baseconfig.DateTimeFormat),
		UpdatedAt: input.UpdatedAt().Format(baseconfig.DateTimeFormat),
	}
}
