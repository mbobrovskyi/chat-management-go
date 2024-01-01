package data_mappers

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/application/common/data_contracts"
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/aggregates/message"
)

func MessageToDTO(msg message.Message) data_contracts.MessageDTO {
	return data_contracts.MessageDTO{
		Id:        msg.Id(),
		Text:      msg.Text(),
		Status:    msg.Status().Uint8(),
		ChatId:    msg.ChatId(),
		CreatedBy: msg.CreatedById(),
		CreatedAt: msg.CreatedAt(),
		UpdatedAt: msg.UpdatedAt(),
	}
}
