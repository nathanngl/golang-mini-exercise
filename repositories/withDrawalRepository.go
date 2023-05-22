package repository

import (
	"database/sql"
	"log"

	model "github.com/nathanngl/golang-mini-exercise.git/models"
)

type WithdrawalRepository struct {
	db *sql.DB
}

func NewWithdrawalRepository(db *sql.DB) *WithdrawalRepository {
	return &WithdrawalRepository{db: db}
}

func (r *WithdrawalRepository) CreateWithdrawal(withdrawal *model.Withdrawal) error {
	stmt, err := r.db.Prepare("INSERT INTO withdrawals (id, withdrawn_by, status, withdrawn_at, amount, reference_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(withdrawal.ID, withdrawal.WithdrawnBy, withdrawal.Status, withdrawal.WithdrawnAt, withdrawal.Amount, withdrawal.ReferenceID)
	if err != nil {
		log.Println("Error executing statement:", err)
		return err
	}

	return nil
}

// Add more repository methods as needed
