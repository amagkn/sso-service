package main

import (
	"github.com/amagkn/sso-service/config"
	"github.com/amagkn/sso-service/internal/app"
	"github.com/amagkn/sso-service/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(cfg.Logger)

	err = app.Run(cfg)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
