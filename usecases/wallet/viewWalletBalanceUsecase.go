package wallet

import (
	"github.com/nathanngl/golang-mini-exercise.git/db"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

// function to call repository to get wallet by owner
func GetWalletByOwner(ownerId string) (*model.Wallet, error) {
	walletRepository := repository.NewWalletRepository(db.GetDB())

	wallet, err := walletRepository.GetWalletByOwner(ownerId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
