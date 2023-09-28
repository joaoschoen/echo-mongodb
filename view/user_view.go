package view

import (
	"API-ECHO-MONGODB/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	g := e.Group("/user")
	g.GET("/:id", controller.GetUser)
	g.GET("/list", controller.GetUserList)
	g.POST("/", controller.PostUser)
	g.PUT("/:id", controller.PutUser)
	g.DELETE("/:id", controller.DeleteUser)
}
