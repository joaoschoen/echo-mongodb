package view

import (
	"API-ECHO-MONGODB/controller/auth"
	"API-ECHO-MONGODB/middleware"
	"API-ECHO-MONGODB/model"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(server *echo.Echo) {
	group := server.Group("/auth")
	group.POST("/login", middleware.CheckBody(auth.Login, model.UnsafeUser{}))
	group.POST("/register", middleware.CheckBody(auth.Register, model.UnsafeUser{}))
}
