package event

import "encoding/json"

type Event struct {
	Type uint8           `json:"type"`
	Data json.RawMessage `json:"data"`
}
