package handler

import (
	"net/http"

	"github.com/Triluong/nc-student/db"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var req db.LoginForm
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	loginResponse, error := db.Login(req)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, error)
	}
	return c.JSON(http.StatusOK, loginResponse)
}
func Register(c echo.Context) error {
	var user db.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := db.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
