package routes

import (
	"github.com/labstack/echo/v4"
	controller "github.com/nathanngl/golang-mini-exercise.git/controllers"
)

func AccountRoute(e *echo.Echo) {
	e.POST("api/v1/init", controller.InitiateAccount)
}
