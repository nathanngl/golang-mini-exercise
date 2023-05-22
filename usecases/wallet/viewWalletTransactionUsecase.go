package wallet

import (
	"errors"
	"sort"
	"time"

	"github.com/nathanngl/golang-mini-exercise.git/db"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

type Transaction struct {
	ID           string    `json:"id"`
	Status       string    `json:"status"`
	TransactedAt time.Time `json:"transacted_at"`
	Type         string    `json:"type"`
	Amount       float64   `json:"amount"`
	ReferenceID  string    `json:"reference_id"`
}

// function to call repository to get all transaction on wallet including deposit and withdrawal
func GetWalletTransactions(ownerId string) ([]Transaction, error) {
	walletRepository := repository.NewWalletRepository(db.GetDB())

	isEnabled, err := walletRepository.IsWalletEnabled(ownerId)
	if err != nil {
		return nil, err
	}

	if !isEnabled {
		return nil, errors.New("wallet is not enabled")
	}

	var transactions []Transaction

	depositRepository := repository.NewDepositRepository(db.GetDB())

	deposits, err := depositRepository.GetDepositsByDepositedBy(ownerId)
	if err != nil {
		return nil, err
	}

	for _, deposit := range deposits {
		transactions = append(transactions, Transaction{
			ID:           deposit.ID,
			Status:       deposit.Status,
			TransactedAt: deposit.DepositedAt,
			Type:         "deposit",
			Amount:       deposit.Amount,
			ReferenceID:  deposit.ReferenceID,
		})
	}

	withdrawalRepository := repository.NewWithdrawalRepository(db.GetDB())

	withdrawals, err := withdrawalRepository.GetWithdrawalsByWithdrawnBy(ownerId)
	if err != nil {
		return nil, err
	}

	for _, withdrawal := range withdrawals {
		transactions = append(transactions, Transaction{
			ID:           withdrawal.ID,
			Status:       withdrawal.Status,
			TransactedAt: withdrawal.WithdrawnAt,
			Type:         "withdrawal",
			Amount:       withdrawal.Amount,
			ReferenceID:  withdrawal.ReferenceID,
		})
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].TransactedAt.After(transactions[j].TransactedAt)
	})

	return transactions, nil
}
