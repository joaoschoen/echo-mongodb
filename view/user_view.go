package view

import (
	"API-ECHO-MONGODB/controller/user"
	"API-ECHO-MONGODB/middleware"
	"API-ECHO-MONGODB/model"

	"github.com/labstack/echo/v4"
)

func UserRoutes(server *echo.Echo) {
	group := server.Group("/user")
	group.GET("/:id", user.GetUser)
	group.GET("/list", user.GetUserList)
	group.PUT("/:id", middleware.CheckBody(user.PutUser, model.UnsafeUser{}))
	group.DELETE("/:id", user.DeleteUser)
}
