package handler

import (
	"net/http"

	"github.com/Triluong/nc-student/db"
	"github.com/labstack/echo/v4"
)

func UpdateStudentByID(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := c.Param("studentID")
	// if err := c.Bind(&student); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	result, err := db.UpdateStudentByID(id, student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}

func GetStudentById(c echo.Context) error {
	id := c.Param("studentID")
	student, err := db.GetStudentByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}

func DeleteStudentById(c echo.Context) error {
	id := c.Param("studentID")
	student, err := db.DeleteStudentById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, student)
}
