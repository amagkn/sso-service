package app

import (
	"github.com/amagkn/sso-service/internal/auth/controller/grpc_router"
	"github.com/amagkn/sso-service/internal/auth/usecase"
)

func AuthDomain(d Dependencies) {
	uc := usecase.New()

	grpc_router.Register(d.GRPCServer, uc)
}
