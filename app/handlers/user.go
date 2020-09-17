package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sshanzel/keep-right/app/dto"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/repository"
)

var _iurepo repository.IUserRepository

//NewUserHandler is the constructor of UserHandlers
func NewUserHandler(iurepo repository.IUserRepository) {
	if _iurepo != nil {
		return
	}

	_iurepo = iurepo
}

// NewUser is the handler for fetching the Users at DB
func NewUser(c echo.Context) error {
	token, err := VerifyToken(c.Request())

	user := new(dto.User)

	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := entities.NewUser(token.UID, user.Name)

	_iurepo.CreateUser(newUser)

	return c.JSON(http.StatusOK, _iurepo.GetUsers())
}

// GetUsers is the handler for fetching the Users at DB
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, _iurepo.GetUsers())
}
