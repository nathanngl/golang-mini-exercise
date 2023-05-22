package model

import "time"

type Wallet struct {
	ID        string     `json:"id"`
	OwnedBy   string     `json:"owned_by"`
	Status    string     `json:"status"`
	EnabledAt *time.Time `json:"enabled_at"`
	Balance   float64    `json:"balance"`
}
