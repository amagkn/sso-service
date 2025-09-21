package app

import "github.com/amagkn/sso-service/internal/auth/controller/grpc_router"

func AuthDomain(d Dependencies) {
	grpc_router.Register(d.gRPCServer)
}
