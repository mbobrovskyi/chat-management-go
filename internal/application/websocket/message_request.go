package websocket

import "github.com/google/uuid"

type MessageRequest struct {
	UUID uuid.UUID `json:"UUID"`
	Text string    `json:"text"`
}
