package api

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/chat"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
)

func ChatToResponse(input chat.Chat) ChatResponse {
	return ChatResponse{
		Id:        input.GetId(),
		Name:      input.GetName(),
		Type:      input.GetType().Uint8(),
		Image:     input.GetImage(),
		CreatedBy: input.GetCreatedBy(),
		CreatedAt: input.GetCreatedAt(),
		UpdatedAt: input.GetUpdatedAt(),
	}
}

func MessageToResponse(input message.Message) MessageResponse {
	return MessageResponse{
		Id:        input.GetId(),
		Text:      input.GetText(),
		Status:    input.GetStatus().Uint8(),
		ChatId:    input.GetChatId(),
		CreatedBy: input.GetCreatedBy(),
		CreatedAt: input.GetCreatedAt(),
		UpdatedAt: input.GetUpdatedAt(),
	}
}
