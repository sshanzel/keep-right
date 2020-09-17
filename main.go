package main

import (
	"github.com/sshanzel/keep-right/app/handlers"
	"github.com/sshanzel/keep-right/app/middlewares"
	"github.com/sshanzel/keep-right/infra/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	handlers.NewUserHandler(repository.NewUserRepository())

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	u := api.Group("/users")

	u.Use(middlewares.Authorize)
	u.GET("", handlers.GetUsers)
	u.GET("/:id", handlers.GetUser)
	u.POST("", handlers.NewUser)

	e.Logger.Fatal(e.Start(":2020"))
}
