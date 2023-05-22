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

func (r *WithdrawalRepository) GetWithdrawalsByWithdrawnBy(withdrawnBy string) ([]*model.Withdrawal, error) {
	query := "SELECT id, withdrawn_by, status, withdrawn_at, amount, reference_id FROM withdrawals WHERE withdrawn_by = ?"

	rows, err := r.db.Query(query, withdrawnBy)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	withdrawals := []*model.Withdrawal{}

	for rows.Next() {
		withdrawal := &model.Withdrawal{}
		err := rows.Scan(&withdrawal.ID, &withdrawal.WithdrawnBy, &withdrawal.Status, &withdrawal.WithdrawnAt, &withdrawal.Amount, &withdrawal.ReferenceID)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		withdrawals = append(withdrawals, withdrawal)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return withdrawals, nil
}
