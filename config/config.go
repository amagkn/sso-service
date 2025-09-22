package config

import (
	"errors"
	"os"
	"time"

	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/grpc_server"
	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/amagkn/sso-service/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	ENV      string        `envconfig:"APP_ENV"         required:"true"`
	Name     string        `envconfig:"APP_NAME"        required:"true"`
	Version  string        `envconfig:"APP_VERSION"     required:"true"`
	TokenTTL time.Duration `envconfig:"APP_TOKEN_TTL"   required:"true"`
}

type Config struct {
	App      App
	Logger   logger.Config
	Postgres postgres.Config
	GRPC     grpc_server.Config
}

func New() (Config, error) {
	var config Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		return config, base_errors.WithPath("godotenv.Load", errors.New("config path is empty"))
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return config, base_errors.WithPath("godotenv.Load", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, base_errors.WithPath("envconfig.Process", err)
	}

	return config, nil
}
