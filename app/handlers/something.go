package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sshanzel/keep-right/app/dto"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/repository"
)

var _isrepo repository.ISomethingRepository

// NewSomethingHandler is the constructor of SomethingHandlers
func NewSomethingHandler(isrepo repository.ISomethingRepository) {
	if isrepo != nil {
		return
	}

	_isrepo = isrepo
}

// NewSomething is the handler for POSTing Something
func NewSomething(c echo.Context) error {
	token, err := GetFirebaseToken(c.Request())

	smthng := new(dto.Something)

	if err = c.Bind(smthng); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := _iurepo.GetUserByUID(token.UID)

	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found!")
	}

	newSmthng := entities.NewSomething(smthng.Title, smthng.Description, smthng.InsideOfID, smthng.SomethingTypeID, user.ID)

	_isrepo.CreateSomething(newSmthng)

	return c.JSON(http.StatusOK, newSmthng)
}

// GetSomethings is the handler for POSTing Something
func GetSomethings(c echo.Context) error {
	token, err := GetFirebaseToken(c.Request())

	user, err := _iurepo.GetUserByUID(token.UID)

	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found!")
	}
	return c.JSON(http.StatusOK, _isrepo.GetSomethingsOfUser(user.ID))
}
