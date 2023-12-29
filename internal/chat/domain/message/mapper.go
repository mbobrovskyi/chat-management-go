package message

import "github.com/mbobrovskyi/chat-management-go/internal/chat/common"

func toDTO(msg Message) common.MessageDTO {
	return common.MessageDTO{
		Id:        msg.GetId(),
		Text:      msg.GetText(),
		Status:    msg.GetStatus().Uint8(),
		ChatId:    msg.GetChatId(),
		CreatedBy: msg.GetCreatedBy(),
		CreatedAt: msg.CreatedAt(),
		UpdatedAt: msg.UpdatedAt(),
	}
}
