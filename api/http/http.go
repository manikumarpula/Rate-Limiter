package http

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	config "rate-limiter/configs"
	"rate-limiter/internal/api_key"
	"rate-limiter/internal/db"
	"rate-limiter/internal/plan"
	"rate-limiter/api/http/routes"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func StartServer(cfg *config.Config) error {
	e := echo.New()

	e.HideBanner = true

	banner, err := os.ReadFile("banner.txt")
	if err != nil {
		log.Fatalf("could not read banner: %v", err)
	}

	fmt.Println(string(banner))

	// middlewares
	slogConfig := slogecho.Config{
		WithSpanID:         true,
		WithTraceID:        true,
		WithRequestBody:    true,
		WithResponseBody:   true,
		WithRequestHeader:  true,
		WithResponseHeader: true,
	}
	e.Use(slogecho.NewWithConfig(cfg.Logger, slogConfig))
	e.Use(middleware.Recover())

	serverCtx := context.Background()

	// Set up CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  cfg.CORS.AllowOrigins,
		AllowMethods:  cfg.CORS.AllowMethods,
		AllowHeaders:  cfg.CORS.AllowHeaders,
		ExposeHeaders: cfg.CORS.ExposeHeaders,
		MaxAge:        cfg.CORS.MaxAge,
	}))

	// db setup
	sqlDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		cfg.Logger.Error("Failed to connect to database", "error", err)
		return err
	}
	defer sqlDB.Close()
	if err := sqlDB.PingContext(serverCtx); err != nil {
		cfg.Logger.Error("Failed to ping database", "error", err)
		return err
	}
	pgxDB, err := pgxpool.New(serverCtx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		cfg.Logger.Error("Failed to create pgx pool", "error", err)
		return err
	}
	defer pgxDB.Close()
	db := db.New(pgxDB)

	// services
	apiKeyService := apiKey.NewApiKeyService(cfg, db)
	planService := plan.NewPlanService(cfg, db)

	// routes
	routes.SetupRateLimiterRoutes(e, *cfg, apiKeyService, planService)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "You are Connected Bro"})
	})

	return e.Start(fmt.Sprintf(":%d", cfg.Database.Port))
}
