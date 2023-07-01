package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Middleware to check API key
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.QueryParam("api_key")

		//Hard code to check user login
		if apiKey != "armin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}
