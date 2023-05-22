package wallet

import (
	"time"

	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"

	"github.com/google/uuid"
)

func CreateWithdraw(ownerId string, referenceId string, amount float64) (*model.Withdrawal, error) {
	withdrawalRepository := repository.NewWithdrawalRepository(db.GetDB())

	newWithdrawal := &model.Withdrawal{
		ID:          uuid.New().String(),
		WithdrawnBy: ownerId,
		Status:      "success",
		WithdrawnAt: time.Now(),
		Amount:      amount,
		ReferenceID: referenceId,
	}

	err := withdrawalRepository.CreateWithdrawal(newWithdrawal)
	if err != nil {
		return nil, err
	}

	walletRepository := repository.NewWalletRepository(db.GetDB())

	walletData, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	newBalance := walletData.Balance - amount

	_, err = walletRepository.UpdateWalletBalance(ownerId, newBalance)
	if err != nil {
		return nil, err
	}

	return newWithdrawal, nil
}
