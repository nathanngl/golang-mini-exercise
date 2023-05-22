package account

import (
	"github.com/google/uuid"
	"github.com/nathanngl/golang-mini-exercise.git/db"
	"github.com/nathanngl/golang-mini-exercise.git/libraries"
	model "github.com/nathanngl/golang-mini-exercise.git/models"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

func InitiateAccount(customerXid string) (string, error) {
	newWallet := &model.Wallet{
		ID:      uuid.New().String(),
		OwnedBy: customerXid,
		Status:  "disabled",
		Balance: 0,
	}

	walletRepository := repository.NewWalletRepository(db.GetDB())

	err := walletRepository.CreateWallet(newWallet)
	if err != nil {
		return "", err
	}

	token := libraries.GenerateWalletToken(customerXid)

	return token, nil
}
