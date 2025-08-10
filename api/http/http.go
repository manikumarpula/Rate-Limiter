package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rate-limiter/configs"

	"github.com/labstack/echo/v4"
)

func StartServer(cfg *config.EnvConfig) {
	e := echo.New()

	e.HideBanner = true

	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatalf("could not read banner: %v", err)
	}
	fmt.Println(string(banner))

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "You are Connected Bro"})
	})
}	
