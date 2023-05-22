package wallet

import (
	"time"

	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"

	"github.com/google/uuid"
)

func CreateDeposit(ownerId string, referenceId string, amount float64) (*model.Deposit, error) {
	depositRepository := repository.NewDepositRepository(db.GetDB())

	newDeposit := &model.Deposit{
		ID:          uuid.New().String(),
		DepositedBy: ownerId,
		Status:      "success",
		DepositedAt: time.Now(),
		Amount:      amount,
		ReferenceID: referenceId,
	}

	err := depositRepository.CreateDeposit(newDeposit)
	if err != nil {
		return nil, err
	}

	walletRepository := repository.NewWalletRepository(db.GetDB())

	walletData, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	newBalance := walletData.Balance + amount

	_, err = walletRepository.UpdateWalletBalance(ownerId, newBalance)
	if err != nil {
		return nil, err
	}

	return newDeposit, nil
}
