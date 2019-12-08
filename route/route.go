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
	g := e.Group("/api/v1/student/staff")
	g.GET("/student/:studentID", handler.GetStudentById)
	g.PATCH("/student/:studentID", handler.UpdateStudent)
}

func Public(e *echo.Echo) {
	g := e.Group("/api/v1/student/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/test", handler.TestPublic)
	g.GET("/student", handler.GetAllStudents)
	g.POST("/student", handler.AddStudent)

}
