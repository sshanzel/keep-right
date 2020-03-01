package main

import (
	"fmt"
	"net/http"

	"github.com/sshanzel/keep-right/infra/db"
	"github.com/sshanzel/keep-right/infra/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":2020"))
}

func hello(c echo.Context) error {
	ctx := db.Connect()
	users := repository.NewUserRepository(ctx)

	fmt.Println(users)

	return c.String(http.StatusOK, "Hello!")
}
