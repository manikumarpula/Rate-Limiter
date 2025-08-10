package routes

import (
	"rate-limiter/api/http/handlers"
	"rate-limiter/api/http/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRateLimiterRoutes(e *echo.Echo) {
	api := e.Group("/api")

	// apply rate limit middleware on all routes
	api.Use(middlewares.AutherizationMiddleware)

	api.GET("/rate-limit", handlers.RateLimitHandler)

}