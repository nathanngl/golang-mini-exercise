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

func (r *DepositRepository) GetDepositsByDepositedBy(depositedBy string) ([]*model.Deposit, error) {
	query := "SELECT id, deposited_by, status, deposited_at, amount, reference_id FROM deposits WHERE deposited_by = ?"

	rows, err := r.db.Query(query, depositedBy)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	deposits := []*model.Deposit{}

	for rows.Next() {
		deposit := &model.Deposit{}
		err := rows.Scan(&deposit.ID, &deposit.DepositedBy, &deposit.Status, &deposit.DepositedAt, &deposit.Amount, &deposit.ReferenceID)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		deposits = append(deposits, deposit)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		return nil, err
	}

	return deposits, nil
}
