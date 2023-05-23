package wallet

import (
	"errors"

	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

func DisableWallet(ownerId string, isDisabled bool) (*model.Wallet, error) {
	walletRepository := repository.NewWalletRepository(db.GetDB())

	isEnabled, err := walletRepository.IsWalletEnabled(ownerId)
	if err != nil {
		return nil, err
	}

	if isEnabled != isDisabled {
		return nil, errors.New("wallet is already disabled")
	}

	_, err = walletRepository.DisableWallet(ownerId)
	if err != nil {
		return nil, err
	}

	updatedWallet, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	return updatedWallet, nil
}
