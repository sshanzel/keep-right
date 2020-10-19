package main

import (
	"net/http"

	"github.com/sshanzel/keep-right/app/handlers"
	"github.com/sshanzel/keep-right/app/middlewares"
	"github.com/sshanzel/keep-right/infra/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createMux() *echo.Echo {
	handlers.NewUserHandler(repository.NewUserRepository())
	handlers.NewSomethingHandler(repository.NewSomethingRepository())

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
	u := api.Group("/users")

	u.Use(middlewares.Authorize)
	u.GET("/me", handlers.GetUser)
	u.GET("/:id", handlers.GetUserByID)
	u.POST("", handlers.NewUser)

	s := api.Group("/somethings")
	s.GET("", handlers.GetSomethings)
	s.POST("", handlers.NewSomething)
	s.PATCH("/:id", handlers.UpdateSomething)

	return e
}
