package grpc_server

import (
	"fmt"
	"net"
	"time"

	"github.com/amagkn/sso-service/pkg/base_errors"
	"github.com/amagkn/sso-service/pkg/logger"
	"google.golang.org/grpc"
)

type Config struct {
	Port    string        `envconfig:"GRPC_PORT"    required:"true"`
	Timeout time.Duration `envconfig:"GRPC_TIMEOUT" required:"true"`
}

type GRPCServer struct {
	Server *grpc.Server
	notify chan error
}

func (g *GRPCServer) start(lis net.Listener) {
	g.notify <- g.Server.Serve(lis)
	close(g.notify)
}

func (g *GRPCServer) Notify() <-chan error {
	return g.notify
}

func New() *GRPCServer {
	return &GRPCServer{
		Server: grpc.NewServer(),
		notify: make(chan error, 1),
	}
}

func (g *GRPCServer) Close() {
	g.Server.GracefulStop()

	logger.Info("GRPC Server closed")
}

func (g *GRPCServer) Run(port string) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return base_errors.WithPath("net.Listen", err)
	}

	go g.start(l)

	logger.Info("GRPC Server started on port: " + port)

	return nil
}
