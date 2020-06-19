package main

import (
	"fmt"
	"net/http"

	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/repository"
	"github.com/sshanzel/update-right/app/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.GetUsers())

	e.Logger.Fatal(e.Start(":2020"))
}

func hello(c echo.Context) error {
	_usersRepo := repository.NewUserRepository()

	newUser := entities.NewUser("test", "sample", "email@sample.com")

	_usersRepo.CreateUser(newUser)

	fmt.Println(*newUser)

	return c.String(http.StatusOK, "Hello!")
}

func helloTest(c echo.Context) error {
	_usersRepo := repository.NewUserRepository()

	newUser := entities.NewUser("test", "sample", "email@sample.com")

	_usersRepo.CreateUser(newUser)

	fmt.Println(*newUser)

	return c.String(http.StatusOK, "Hello! Test")
}
