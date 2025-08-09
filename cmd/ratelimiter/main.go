package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"ratelimiter/internal/config"
)

func main() {
	log.Println("Rate Limiter service starting...")

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("could not load configuration: %v", err)
	}

	log.Printf("Loaded configuration: server=%s:%d, algo=%s",
		cfg.Server.Host, cfg.Server.Port, cfg.RateLimiter.Algorithm)

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
