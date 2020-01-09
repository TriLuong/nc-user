package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Triluong/nc-user/config"
	"github.com/Triluong/nc-user/db"
	MyMiddleware "github.com/Triluong/nc-user/middleware"
	"github.com/Triluong/nc-user/route"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	db.Init()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(MyMiddleware.SimpleLogger())
	route.All(e)
	log.Println(e.Start(":9090"))
}
