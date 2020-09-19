package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sshanzel/keep-right/app/handlers"
	"github.com/sshanzel/keep-right/infra/services"
)

// Authorize is a middleware to allow access to protected endpoints
func Authorize(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		token := handlers.ExtractToken(c.Request())

		_, err := services.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())

			return nil
		}

		if err := next(c); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}
