package customMiddleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nathanngl/golang-mini-exercise.git/db"
	"github.com/nathanngl/golang-mini-exercise.git/libraries"
	repository "github.com/nathanngl/golang-mini-exercise.git/repositories"
)

func TokenAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customHeaderValue := c.Request().Header.Get("Authorization")
		if customHeaderValue == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Unauthorized",
			})
		}

		// it should be more authorization checking here, but lack of information about it :/
		// this is just assumptions
		walletRepository := repository.NewWalletRepository(db.GetDB())
		wallets, err := walletRepository.GetAllWallet()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
		}

		// map if wallet owner_by encrypted to sha1 is equal to customHeaderValue
		// if not equal, return unauthorized
		for _, wallet := range wallets {
			encryptedOwnerId := libraries.GenerateWalletToken(wallet.OwnedBy)
			token := strings.Split(customHeaderValue, " ")[1]
			if encryptedOwnerId == token {
				c.Set("ownerId", wallet.OwnedBy)
				return next(c)
			}
		}

		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
	}
}
