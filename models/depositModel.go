package model

import "time"

type Deposit struct {
	ID          string    `json:"id"`
	DepositedBy string    `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceID string    `json:"reference_id"`
}
