package websocket

import "github.com/google/uuid"

type CreateMessageRequest struct {
	UUID uuid.UUID `json:"UUID"`
	Text string    `json:"text"`
}
