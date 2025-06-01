package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	// TestEnv    string `envconfig:"TESTENV" required:"true"`
	DBHost     string `envconfig:"POSTGRES_HOST" required:"true"`
	DBPort     string `envconfig:"POSTGRES_PORT" required:"true"`
	DBUser     string `envconfig:"POSTGRES_USER" required:"true"`
	DBPassword string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DBName     string `envconfig:"POSTGRES_DB" required:"true"`
	ServerPort string `envconfig:"PRODUCT_SERVER_PORT" required:"true"`
	APIKey     string `envconfig:"PUBG_API_KEY" required:"true"`
	APIMaxRPM  string `envconfig:"REQUESTS_PER_MINUTE" required:"true"`
}

// NewEnv initializes the EnvConfig by loading the environment variables.
func NewEnv() *EnvConfig {
	var env EnvConfig
	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("Failed to load env vars: %s", err)
	}
	return &env
}
