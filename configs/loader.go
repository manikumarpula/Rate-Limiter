package config

import (
	"os"
)

func LoadConfigFromEnv() (EnvConfig, error) {
	var cfg EnvConfig

	cfg.SecretKey = os.Getenv("SECRET_KEY")

	return cfg, nil
}
