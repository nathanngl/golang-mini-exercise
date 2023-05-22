package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nathanngl/golang-mini-exercise.git/libraries"
	usecase "github.com/nathanngl/golang-mini-exercise.git/usecases/account"
)

func InitiateAccount(c echo.Context) error {
	body, err := c.FormParams()
	if err != nil {
		fmt.Println("Error serializing response:", err)
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError("Internal server error"))
	}

	if body.Get("customer_xid") == "" {
		data := map[string]interface{}{
			"error": map[string]interface{}{
				"customer_xid": "Missing data for required field.",
			},
		}
		return c.JSON(http.StatusBadRequest, libraries.ResponseFail(data))
	}

	token, err := usecase.InitiateAccount(body.Get("customer_xid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libraries.ResponseError("Internal server error"))
	}

	result := map[string]interface{}{
		"token": token,
	}

	return c.JSON(http.StatusOK, libraries.ResponseOK(result))
}
