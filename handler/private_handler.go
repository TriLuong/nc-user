package handler

import (
	"log"
	"net/http"

	"github.com/Triluong/nc-student/db"
	"github.com/labstack/echo/v4"
)

func UpdateUserByID(c echo.Context) error {
	var user db.UserUpdate
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := c.Param("userID")
	if user.Password != "" {
		password, error := db.HashPassword(user.Password)
		if error != nil {
			log.Println(error)
			return error
		}
		user.Password = password
	}
	result, err := db.UpdateUserByID(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}
