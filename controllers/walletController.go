package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nathanngl/golang-mini-exercise.git/libraries"
	usecase "github.com/nathanngl/golang-mini-exercise.git/usecases/wallet"
)

// function to enable wallet
func EnableWallet(c echo.Context) error {
	ownerId := c.Get("ownerId").(string)
	log.Println(ownerId)
	enableWallet, err := usecase.EnableWallet(ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError(err.Error()))
	}

	return c.JSON(http.StatusOK, libraries.ResponseOK(enableWallet))
}

func ViewWalletBalance(c echo.Context) error {
	ownerId := c.Get("ownerId").(string)
	log.Println(ownerId)
	wallet, err := usecase.GetWalletByOwner(ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError(err.Error()))
	}

	result := map[string]interface{}{
		"wallet": wallet,
	}

	return c.JSON(http.StatusOK, libraries.ResponseOK(result))
}

func TopUpWallet(c echo.Context) error {
	ownerId := c.Get("ownerId").(string)
	referenceId := c.FormValue("reference_id")
	amount, _ := strconv.ParseFloat(c.FormValue("amount"), 64)

	deposit, err := usecase.CreateDeposit(ownerId, referenceId, amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError(err.Error()))
	}

	result := map[string]interface{}{
		"deposit": deposit,
	}

	return c.JSON(http.StatusOK, libraries.ResponseOK(result))
}

func WithdrawWallet(c echo.Context) error {
	ownerId := c.Get("ownerId").(string)
	referenceId := c.FormValue("reference_id")
	amount, _ := strconv.ParseFloat(c.FormValue("amount"), 64)

	withdraw, err := usecase.CreateWithdraw(ownerId, referenceId, amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError(err.Error()))
	}

	result := map[string]interface{}{
		"withdraw": withdraw,
	}

	return c.JSON(http.StatusOK, libraries.ResponseOK(result))
}
