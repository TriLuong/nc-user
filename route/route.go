package route

import (
	"github.com/Triluong/nc-student/handler"
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

}

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/student/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestPublic)

}
