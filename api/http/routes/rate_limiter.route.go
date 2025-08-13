package routes

import (
	"rate-limiter/api/http/handlers"
	"rate-limiter/api/http/middlewares"
	config "rate-limiter/configs"

	"github.com/labstack/echo/v4"
)

func SetupRateLimiterRoutes(e *echo.Echo, config config.EnvConfig) {
	api := e.Group("/api")

	// apply rate limit middleware on all routes
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return middlewares.AutherizationMiddleware(next, &config)
	})

	api.GET("/rate-limit", handlers.RateLimitHandler)

}
