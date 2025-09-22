package app

import (
	"time"

	"github.com/amagkn/sso-service/internal/auth/adapters/postgres"
	"github.com/amagkn/sso-service/internal/auth/controller/grpc_router"
	"github.com/amagkn/sso-service/internal/auth/usecase"
)

func AuthDomain(d Dependencies, tokenTTL time.Duration) {
	uc := usecase.New(postgres.New(d.Postgres), tokenTTL)

	grpc_router.Register(d.GRPCServer, uc)
}
