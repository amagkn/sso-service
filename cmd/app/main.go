package main

import (
	"context"

	"github.com/amagkn/sso-service/config"
	"github.com/amagkn/sso-service/internal/app"
	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/amagkn/sso-service/pkg/validation"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "config.New")
	}

	logger.Init(cfg.Logger)
	validation.Init()

	err = app.Run(ctx, cfg)
	if err != nil {
		logger.Fatal(err, "app.Run")
	}

	logger.Info("App stopped!")
}
