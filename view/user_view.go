package view

import (
	"API-ECHO-MONGODB/controller/user"

	"github.com/labstack/echo/v4"
)

func UserRoutes(server *echo.Echo) {
	group := server.Group("/user")
	group.POST("", user.PostUser)
	group.GET("/:id", user.GetUser)
	group.GET("/list", user.GetUserList)
	group.PUT("/:id", user.PutUser)
	group.DELETE("/:id", user.DeleteUser)
}
