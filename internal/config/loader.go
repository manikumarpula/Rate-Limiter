package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (Config, error) {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return cfg, nil
}
