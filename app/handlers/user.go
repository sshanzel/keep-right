package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sshanzel/keep-right/infra/repository"
)

// GetUsers is the handler for fetching the Users at DB
func GetUsers(c echo.Context) error {
	var iurepo repository.IUserRepository
	iurepo = repository.NewUserRepository()

	return c.JSON(http.StatusOK, iurepo.GetUsers())
}
