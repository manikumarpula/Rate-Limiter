package routes

import (
	"rate-limiter/api/http/handlers"
	"rate-limiter/api/http/middlewares"
	config "rate-limiter/configs"
	"rate-limiter/internal/api_key"
	"rate-limiter/internal/plan"

	"github.com/labstack/echo/v4"
)

func SetupRateLimiterRoutes(e *echo.Echo, config config.Config, apiKeyService *apiKey.ApiKeyService, planService *plan.PlanService) {
	api := e.Group("/api")

	// apply rate limit middleware on all routes
	api.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return middlewares.AutherizationMiddleware(next, &config)
	})

	api.GET("/rate-limit", handlers.RateLimitHandler)

}
