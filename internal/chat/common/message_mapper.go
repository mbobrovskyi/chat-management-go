package common

import (
	"github.com/mbobrovskyi/chat-management-go/internal/chat/domain/message"
)

func MessageToDTO(msg message.Message) MessageDTO {
	return MessageDTO{
		Id:        msg.GetId(),
		Text:      msg.GetText(),
		Status:    msg.GetStatus().Uint8(),
		ChatId:    msg.GetChatId(),
		CreatedBy: msg.GetCreatedBy(),
		CreatedAt: msg.CreatedAt(),
		UpdatedAt: msg.UpdatedAt(),
	}
}

func MessageFromDTO(msg MessageDTO) message.Message {
	return message.Create(msg.Id, msg.Text, message.NewMessageStatus(msg.Status), msg.ChatId, msg.CreatedBy, msg.CreatedAt, msg.UpdatedAt)
}
