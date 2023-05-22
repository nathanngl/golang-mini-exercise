package model

import "time"

type Withdrawal struct {
	ID          string    `json:"id"`
	WithdrawnBy string    `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}
