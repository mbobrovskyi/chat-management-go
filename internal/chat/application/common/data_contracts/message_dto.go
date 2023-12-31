package data_contracts

import (
	"time"
)

type MessageDTO struct {
	Id        uint64    `json:"id"`
	Text      string    `json:"text"`
	Status    uint8     `json:"status"`
	ChatId    uint64    `json:"chatId"`
	CreatedBy uint64    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
