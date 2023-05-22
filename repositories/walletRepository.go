package repository

import (
	"database/sql"
	"log"
	"time"

	model "github.com/nathanngl/golang-mini-exercise.git/models"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) CreateWallet(wallet *model.Wallet) error {
	stmt, err := r.db.Prepare("INSERT INTO wallets (id, owned_by, status, balance) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(wallet.ID, wallet.OwnedBy, wallet.Status, wallet.Balance)
	if err != nil {
		log.Println("Error executing statement:", err)
		return err
	}

	return nil
}

func (r *WalletRepository) EnableWallet(ownerId string) (map[string]interface{}, error) {
	stmt, err := r.db.Prepare("UPDATE wallets SET status = ?, enabled_at = ? WHERE owned_by = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec("enabled", time.Now(), ownerId)
	if err != nil {
		log.Println("Error executing statement:", err)
		return nil, err
	}

	return map[string]interface{}{
		"status": "enabled",
	}, nil
}

func (r *WalletRepository) GetWalletByOwner(ownerId string) (*model.Wallet, error) {
	stmt, err := r.db.Prepare("SELECT * FROM wallets WHERE owned_by = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	var wallet model.Wallet
	err = stmt.QueryRow(ownerId).Scan(&wallet.ID, &wallet.OwnedBy, &wallet.Status, &wallet.EnabledAt, &wallet.Balance)
	if err != nil {
		log.Println("Error scanning rows:", err)
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) GetAllWallet() ([]*model.Wallet, error) {
	stmt, err := r.db.Prepare("SELECT * FROM wallets")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("Error executing statement:", err)
		return nil, err
	}
	defer rows.Close()

	var wallets []*model.Wallet
	for rows.Next() {
		var wallet model.Wallet
		err := rows.Scan(&wallet.ID, &wallet.OwnedBy, &wallet.Status, &wallet.EnabledAt, &wallet.Balance)
		if err != nil {
			log.Println("Error scanning rows:", err)
			return nil, err
		}

		wallets = append(wallets, &wallet)
	}

	return wallets, nil
}

func (r *WalletRepository) UpdateWalletBalance(ownerId string, amount float64) (map[string]interface{}, error) {
	stmt, err := r.db.Prepare("UPDATE wallets SET balance = ? WHERE owned_by = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(amount, ownerId)
	if err != nil {
		log.Println("Error executing statement:", err)
		return nil, err
	}

	return map[string]interface{}{
		"balance": amount,
	}, nil
}
