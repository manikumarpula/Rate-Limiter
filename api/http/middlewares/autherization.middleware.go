package middlewares

import (
	"errors"
	"net/http"
	"rate-limiter/configs"
	"strings"
	"rate-limiter/internal/auth"

	"github.com/labstack/echo/v4"
)

func AutherizationMiddleware(next echo.HandlerFunc,config *config.EnvConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check for presence of Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, errors.New("authorization header is missing"))
		}

		// Extract token from Bearer format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, errors.New("invalid authorization format"))
		}

		// Trim the Bearer prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, errors.New("jwt token is missing"))
		}

		// Validate token
		jwtService := auth.NewJwtService(config.SecretKey, config.SecretKey)

		_, err := jwtService.DecodeToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
		}

		return next(c)
	}
}
