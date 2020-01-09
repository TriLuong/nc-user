package route

import (
	"github.com/Triluong/nc-user/handler"
	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {
	g := e.Group("/api/v1/user/private")
	g.PATCH("/user/:userID", handler.UpdateUserByID)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/user/public")
	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
	// g.POST("/register", handler.Register)
}
