package handlers

import (
	"net/http"

	"github.com/google/uuid"
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
	token, err := GetFirebaseToken(c.Request())

	user := new(dto.User)

	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newUser := entities.NewUser(token.UID, user.Name)

	_iurepo.CreateUser(newUser)

	return c.JSON(http.StatusOK, _iurepo.GetUsers())
}

// GetUsers returns all users in DB
func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, _iurepo.GetUsers())
}

// GetUser returns user with the specified id (PK)
func GetUser(c echo.Context) error {
	token, _ := GetFirebaseToken(c.Request())

	user, err := _iurepo.GetUserByUID(token.UID)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// GetUserByID returns user with the specified id (PK)
func GetUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	user := _iurepo.GetUser(id)

	if user == nil {
		return c.JSON(http.StatusNotFound, "User doesn't exists")
	}

	token, err := GetFirebaseToken(c.Request())

	if token.UID != user.UID {
		return c.JSON(http.StatusForbidden, "Not Allowed")
	}

	return c.JSON(http.StatusOK, user)
}
