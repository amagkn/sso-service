package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/amagkn/sso-service/config"
	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/grpc_server"
	"github.com/amagkn/sso-service/pkg/logger"
	"github.com/amagkn/sso-service/pkg/postgres"
)

type Dependencies struct {
	GRPCServer *grpc_server.GRPCServer
	Postgres   *postgres.Pool
}

func Run(ctx context.Context, cfg config.Config) (err error) {
	var deps Dependencies

	deps.Postgres, err = postgres.New(ctx, cfg.Postgres)
	if err != nil {
		return base_errors.WithPath("postgres.New", err)
	}
	defer deps.Postgres.Close()

	gRPCServer := grpc_server.New()
	deps.GRPCServer = gRPCServer

	AuthDomain(deps)

	err = gRPCServer.Run(cfg.GRPC.Port)
	if err != nil {
		return base_errors.WithPath("grpc_server.Run", err)
	}
	defer gRPCServer.Close()

	waiting(gRPCServer)

	return nil
}

func waiting(gRPCServer *grpc_server.GRPCServer) {
	logger.Info("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		logger.Info("App got signal: " + i.String())
	case err := <-gRPCServer.Notify():
		logger.Error(err, "App got notify: gRPCServer.Notify")
	}

	logger.Info("App is stopping")
}
