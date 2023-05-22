package wallet

import (
	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

// function to call repository to update wallet status
func EnableWallet(ownerId string) (*model.Wallet, error) {
	walletRepository := repository.NewWalletRepository(db.GetDB())

	_, err := walletRepository.EnableWallet(ownerId)
	if err != nil {
		return nil, err
	}

	updatedWallet, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	return updatedWallet, nil
}
