package http

import "github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/message"

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
