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
