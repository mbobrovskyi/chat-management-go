package api

import (
	"time"
)

type ChatResponse struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Type      uint8     `json:"type"`
	Image     string    `json:"image"`
	CreatedBy uint64    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MessageResponse struct {
	Id        uint64    `json:"id"`
	Text      string    `json:"text"`
	Status    uint8     `json:"status"`
	ChatId    uint64    `json:"chatId"`
	CreatedBy uint64    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
