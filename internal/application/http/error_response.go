package http

type ErrorResponse struct {
	Timestamp string         `json:"timestamp,omitempty"`
	Code      string         `json:"code,omitempty"`
	Message   string         `json:"message,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}
