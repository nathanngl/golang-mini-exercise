package wallet

import (
	"errors"

	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

// function to call repository to update wallet status
func EnableWallet(ownerId string) (*model.Wallet, error) {
	walletRepository := repository.NewWalletRepository(db.GetDB())

	isEnabled, err := walletRepository.IsWalletEnabled(ownerId)
	if err != nil {
		return nil, err
	}

	if isEnabled {
		return nil, errors.New("wallet is already enabled")
	}

	_, err = walletRepository.EnableWallet(ownerId)
	if err != nil {
		return nil, err
	}

	updatedWallet, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	return updatedWallet, nil
}
