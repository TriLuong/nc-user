package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Triluong/nc-student/route"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	route.All(e)

	log.Println(e.Start(":9090"))
}
