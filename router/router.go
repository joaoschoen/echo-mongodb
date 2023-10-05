package router

import (
	"API-ECHO-MONGODB/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {

	view.UserRoutes(server)
	view.AuthRoutes(server)
}
