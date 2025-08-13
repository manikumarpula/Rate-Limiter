package config

import "log/slog"

type Config struct {
	Logger   *slog.Logger
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type DatabaseConfig struct {
	Host         string
	Port         string
	Name         string
	Username     string
	Password     string
	MigrationDir string
}

type JWTConfig struct {
	SecretKey            string
	Issuer               string
	AccessTokenLifetime  int
	RefreshTokenLifetime int
}

type CORSConfig struct {
	AllowOrigins  []string
	AllowMethods  []string
	AllowHeaders  []string
	ExposeHeaders []string
	MaxAge        int
}
