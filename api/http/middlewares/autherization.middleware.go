package middlewares

import (
	"net/http"
	"rate-limiter/api/http/handlers/dtos"

	"github.com/labstack/echo/v4"
)

func AutherizationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.JSON(http.StatusUnauthorized, dtos.BaseResponse{
					Status:  "error",
					Message: "Unauthorized",
					Error:   "Unauthorized",
				})
			}
			return next(c)
		}
	}
}
