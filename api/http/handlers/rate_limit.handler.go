package handlers

import (
	"net/http"
	"rate-limiter/api/http/handlers/dtos"

	"github.com/labstack/echo/v4"
)

func RateLimitHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, dtos.BaseResponse{
		Status:  "success",
		Message: "Rate limit",
		Data:    nil,
	})
}