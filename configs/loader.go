package config

import (
	"io"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadConfigFromEnv() (*Config, error) {
	// Initialize the logger
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}

	// Create a multi-writer: file + stdout
	multiWriter := io.MultiWriter(os.Stdout, file)
	logger := slog.New(slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Load .env variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse numeric values
	corsMaxAge, err := strconv.Atoi(os.Getenv("CORS_MAX_AGE"))
	if err != nil {
		log.Fatalf("Invalid CORS_MAX_AGE value: %v", err)
	}

	accessTokenLifetime, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalf("Invalid ACCESS_TOKEN_LIFETIME value: %v", err)
	}

	refreshTokenLifetime, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME"))
	if err != nil {
		log.Fatalf("Invalid REFRESH_TOKEN_LIFETIME value: %v", err)
	}

	config := &Config{
		Logger: logger,
		Database: DatabaseConfig{
			Host:         os.Getenv("DB_HOST"),
			Port:         os.Getenv("DB_PORT"),
			Name:         os.Getenv("DB_NAME"),
			Username:     os.Getenv("DB_USERNAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			MigrationDir: os.Getenv("DB_MIGRATION_DIR"),
		},
		JWT: JWTConfig{
			SecretKey:            os.Getenv("JWT_SECRET_KEY"),
			Issuer:               os.Getenv("JWT_ISSUER"),
			AccessTokenLifetime:  accessTokenLifetime,
			RefreshTokenLifetime: refreshTokenLifetime,
		},
		CORS: CORSConfig{
			AllowOrigins:  strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"), ","),
			AllowMethods:  strings.Split(os.Getenv("CORS_ALLOW_METHODS"), ","),
			AllowHeaders:  strings.Split(os.Getenv("CORS_ALLOW_HEADERS"), ","),
			ExposeHeaders: strings.Split(os.Getenv("CORS_EXPOSE_HEADERS"), ","),
			MaxAge:        corsMaxAge,
		},
	}

	// Validate required fields
	if config.Database.Host == "" {
		log.Fatal("DB_HOST is not set in the environment variables")
	}
	if config.Database.Port == "" {
		log.Fatal("DB_PORT is not set in the environment variables")
	}
	if config.Database.Name == "" {
		log.Fatal("DB_NAME is not set in the environment variables")
	}
	if config.Database.Username == "" {
		log.Fatal("DB_USERNAME is not set in the environment variables")
	}
	if config.Database.Password == "" {
		log.Fatal("DB_PASSWORD is not set in the environment variables")
	}
	if config.Database.MigrationDir == "" {
		log.Fatal("DB_MIGRATION_DIR is not set in the environment variables")
	}
	if config.JWT.SecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in the environment variables")
	}
	if config.JWT.Issuer == "" {
		log.Fatal("JWT_ISSUER is not set in the environment variables")
	}

	// Log the loaded configuration
	logger.Info("Configuration loaded successfully",
		"db_host", config.Database.Host,
		"db_port", config.Database.Port,
		"db_name", config.Database.Name,
		"db_username", "******",
		"db_password", "******",
		"db_migration_dir", config.Database.MigrationDir,
		"jwt_issuer", config.JWT.Issuer,
		"access_token_lifetime", config.JWT.AccessTokenLifetime,
		"refresh_token_lifetime", config.JWT.RefreshTokenLifetime,
	)

	return config, nil
}
