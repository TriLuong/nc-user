package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Triluong/nc-student/config"
	"github.com/Triluong/nc-student/db"
	"github.com/Triluong/nc-student/route"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	db.Init()
	e := echo.New()
	e.Use(middleware.Recover())
	route.All(e)
	log.Println(e.Start(":9090"))
}
