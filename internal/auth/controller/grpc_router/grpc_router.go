package grpc_router

import (
	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
	"github.com/amagkn/sso-service/internal/auth/controller/grpc_router/ver1"
	"github.com/amagkn/sso-service/internal/auth/usecase"
	"github.com/amagkn/sso-service/pkg/grpc_server"
)

func Register(gRPC *grpc_server.GRPCServer, uc *usecase.UseCase) {
	v1 := ver1.New(uc)

	ssov1.RegisterAuthServer(gRPC.Server, v1)
}
