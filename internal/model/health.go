package model

type Health struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
}
