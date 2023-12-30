package services

import (
	"github.com/mbobrovskyi/chat-management-go/internal/common"
	"github.com/mbobrovskyi/chat-management-go/internal/domain/aggregates/message"
)

func MessageToDTO(msg message.Message) common.MessageDTO {
	return common.MessageDTO{
		Id:        msg.Id(),
		Text:      msg.Text(),
		Status:    msg.Status().Uint8(),
		ChatId:    msg.ChatId(),
		CreatedBy: msg.CreatedById(),
		CreatedAt: msg.CreatedAt(),
		UpdatedAt: msg.UpdatedAt(),
	}
}
