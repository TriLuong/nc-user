package handler

import (
	"github.com/Triluong/nc-student/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestPublic(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "TestPublic")
}
