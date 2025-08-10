package routes

import (
	"net/http"
	"rate-limiter/api/http/middlewares"
	"rate-limiter/api/http/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "You are Connected Bro"})
	})

	api := e.Group("/api")

	// apply rate limit middleware on all routes
	api.Use(middlewares.AutherizationMiddleware())

	api.GET("/rate-limit", handlers.RateLimitHandler)

}