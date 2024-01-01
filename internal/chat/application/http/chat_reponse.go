package http

type ChatResponse struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        uint8  `json:"type"`
	Image       string `json:"image"`
	CreatedBy   uint64 `json:"createdBy"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
