package middlewares

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sshanzel/keep-right/app/handlers"
)

// Authorize is a middleware to allow access to protected endpoints
func Authorize(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		_, err := handlers.VerifyToken(c.Request())

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())

			return nil
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
