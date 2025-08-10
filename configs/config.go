package config

// Config represents the application configuration
type Config struct {
	Server struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		Timeout string `mapstructure:"timeout"`
	} `mapstructure:"server"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

type EnvConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}
