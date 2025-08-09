package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"rate-limiter/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Rate Limiter service starting...")

	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("could not load configuration: %v", err)
	}

	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatalf("could not read banner: %v", err)
	}
	fmt.Println(string(banner))

	e := echo.New()
	e.HideBanner = true
	
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
