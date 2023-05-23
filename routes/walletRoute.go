package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/nathanngl/golang-mini-exercise.git/controllers"
	customMiddleware "github.com/nathanngl/golang-mini-exercise.git/middleware"
)

func WalletRoute(e *echo.Echo) {
	wallet := e.Group("/api/v1/wallet")
	wallet.Use(customMiddleware.TokenAuth)
	wallet.POST("", controller.EnableWallet)
	wallet.GET("", controller.ViewWalletBalance)
	wallet.POST("/deposits", controller.TopUpWallet)
	wallet.POST("/withdrawals", controller.WithdrawWallet)
	wallet.GET("/transactions", controller.ViewWalletTransactions)
	wallet.PATCH("", controller.DisableWallet)
}
