package repository

import (
	"database/sql"
	"log"

	model "github.com/nathanngl/golang-mini-exercise.git/models"
)

type DepositRepository struct {
	db *sql.DB
}

func NewDepositRepository(db *sql.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (r *DepositRepository) CreateDeposit(deposit *model.Deposit) error {
	stmt, err := r.db.Prepare("INSERT INTO deposits (id, deposited_by, status, deposited_at, amount, reference_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(deposit.ID, deposit.DepositedBy, deposit.Status, deposit.DepositedAt, deposit.Amount, deposit.ReferenceID)
	if err != nil {
		log.Println("Error executing statement:", err)
		return err
	}

	return nil
}

// Add more repository methods as needed
