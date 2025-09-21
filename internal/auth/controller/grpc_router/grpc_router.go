package grpc_router

import (
	"context"

	ssov1 "github.com/amagkn/sso-protos/gen/go/sso"
	"github.com/amagkn/sso-service/pkg/grpc_server"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc_server.GRPCServer) {
	ssov1.RegisterAuthServer(gRPC.Server, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	return &ssov1.LoginResponse{Token: req.GetEmail()}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("not implement")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("not implement")
}
